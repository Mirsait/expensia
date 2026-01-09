package cmd

import (
	"fmt"

	"github.com/Mirsait/expensia/models"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// required flag
var Id int

func init() {
	deleteCmd.Flags().IntVar(
		&Id,
		"id",
		0,
		"expense id")
	deleteCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete by ID",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		datafile := viper.GetString("data")
		if data, err := storage.Load(datafile); err == nil {
			var newData []models.Expense
			for j, item := range data {
				if item.Id == Id {
					newData = append(data[:j], data[j+1:]...)
					break
				}
			}
			if newData == nil {
				fmt.Println("Item not found")
			} else if err = storage.Save(datafile, newData); err == nil {
				fmt.Println("Expense deleted successfully")
			} else {
				fmt.Println("Saving error:", err.Error())
			}
		} else {
			fmt.Println("Loading error:", err.Error())
		}
	},
}
