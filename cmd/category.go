package cmd

import (
	"fmt"

	"github.com/Mirsait/expensia/fp"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(categoryCmd)
}

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "Show expenses category list",
	Run: func(cmd *cobra.Command, args []string) {
		datafile := viper.GetString("data")
		if data, err := storage.Load(datafile); err == nil {
			cats := fp.Reduce(data,
				make([]string, 0),
				func(acc []string, x Expense) []string {
					return append(acc, x.Category)
				})
			cats = fp.Distinct(cats)
			for _, cat := range cats {
				fmt.Println(cat)
			}
		} else {
			fmt.Println("Loading error: ", err.Error())
		}
	},
}
