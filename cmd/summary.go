package cmd

import (
	"fmt"
	"time"

	"github.com/Mirsait/expensia/models"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Expense = models.Expense

var monthValue int

func init() {
	summaryCmd.Flags().IntVar(
		&monthValue,
		"month",
		0,
		"filter by month")
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show summary for the current year or month in the current year",
	Run: func(cmd *cobra.Command, args []string) {
		if monthValue < 0 || monthValue > 12 {
			fmt.Println("Input error: The month must be between 1 and 12 inclusive.")
			return
		}
		datafile := viper.GetString("data")
		if data, err := storage.Load(datafile); err == nil {
			total := 0
			currentYear := time.Now().Year()
			if monthValue > 0 {
				total = expensesInMonth(data, currentYear, monthValue)
			} else {
				total = expensesInCurrentYear(data, currentYear)
			}

			fmt.Printf("Total expenses: $%d\n", total)
		} else {
			fmt.Println("Loading error:", err.Error())
		}
	},
}

func expensesInCurrentYear(data []Expense, currentYear int) int {
	return filterData(data, func(x Expense) bool {
		return x.CreatedAt.Year() == currentYear
	})
}
func expensesInMonth(data []Expense, currentYear, month int) int {
	return filterData(data, func(x Expense) bool {
		return x.CreatedAt.Year() == currentYear &&
			x.CreatedAt.Month() == time.Month(month)
	})
}

func filterData(data []Expense, pred func(Expense) bool) int {
	total := 0
	for _, item := range data {
		if pred(item) {
			total += item.Amount
		}
	}
	return total
}
