package cmd

import (
	"fmt"

	"github.com/Mirsait/expensia/models"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// local flags will only be used with the `add` command
var Description string
var Amount int

func init() {
	addCmd.Flags().StringVar(
		&Description,
		"description",
		"",
		"an expense description")
	addCmd.MarkFlagRequired("description")
	addCmd.Flags().IntVar(
		&Amount,
		"amount",
		0,
		"an expense amount")
	addCmd.MarkFlagRequired("amount")
	addCmd.MarkFlagsRequiredTogether("description", "amount")
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if err  := validData(Description, Amount); err != nil {
			fmt.Println("Input error:", err.Error())
			return
		}
		datafile := viper.GetString("data")
		if data, err := storage.Load(datafile); err == nil {
			nextId := 1
			for _, item := range data {
				if item.Id >= nextId {
					nextId = item.Id + 1
				}
			}
			newExp := models.NewExpense(nextId, Description, Amount)
			newData := append(data, newExp)
			if saveErr := storage.Save(datafile, newData); saveErr == nil {
				fmt.Printf("Expense added successfully (ID: %d)\n", nextId)
			} else {
				fmt.Println("Saving error: ", err.Error())
			}
		} else {
			fmt.Println("Loading error: ", err.Error())
		}
	},
}

func validData(description string, amount int) error {
	if len(description) == 0 {
		return fmt.Errorf("Description cannot be empty")
	}
	if amount < 0 {
		return fmt.Errorf("Amount cannot be negative")
	}
	return nil
}
