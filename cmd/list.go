package cmd

import (
	"fmt"

	"github.com/Mirsait/expensia/fp"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	listCmd.Flags().StringVarP(
		&Category,
		CategoryFlag,
		"c",
		defaultCategory,
		"Filter expenses by category")
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show expenses list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		datafile := viper.GetString("data")
		if data, err := storage.Load(datafile); err == nil {
			if Category == defaultCategory {
				printData(data)
			} else {
				filtered := fp.Filter(data, func(x Expense) bool {
					return x.Category == Category
				})
				printData(filtered)
			}
		}
	},
}

func printData(data []Expense) {
	fmt.Printf("%4s %12s %15s %30s %6s\n",
		"ID", "Date", "Category", "Description", "Amount")
	for _, e := range data {
		prettyDate := e.CreatedAt.Format("02-01-2006")
		description := cutDescription(e.Description)
		fmt.Printf("%4d %12s %15s %30s $%d\n",
			e.Id, prettyDate, e.Category, description, e.Amount)

	}
}

func cutDescription(text string) string {
	if len(text) > 30 {
		return text[:30]
	}
	return text
}
