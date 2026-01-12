package cmd

import (
	"fmt"
	"time"

	"github.com/Mirsait/expensia/fp"
	"github.com/Mirsait/expensia/models"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Expense = models.Expense

var monthValue int

func init() {
	summaryCmd.Flags().IntVarP(
		&monthValue,
		MonthFlag,
		"m",
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
			currentYear := time.Now().Year()
			if monthValue > 0 {
				total := expensesInMonth(data, currentYear, monthValue)
				monthName := time.Month(monthValue).String()
				fmt.Printf("Total expenses for %s: $%d\n", monthName, total)
			} else {
				total := expensesInCurrentYear(data, currentYear)
				fmt.Printf("Total expenses: $%d\n", total)
			}
		} else {
			fmt.Println("Loading error:", err.Error())
		}
	},
}

func expensesInCurrentYear(data []Expense, currentYear int) int {
	return fp.ReduceWithFilter(
		data,
		0,
		func(acc int, x Expense) int {
			return acc + x.Amount
		},
		func(x Expense) bool {
			return x.CreatedAt.Year() == currentYear
		})
}

func expensesInMonth(data []Expense, currentYear, month int) int {
	return fp.ReduceWithFilter(
		data,
		0,
		func(acc int, x Expense) int {
			return acc + x.Amount
		},
		func(x Expense) bool {
			date := x.CreatedAt
			return date.Year() == currentYear &&
				date.Month() == time.Month(month)
		})
}
