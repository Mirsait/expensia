package cmd

import (
	"fmt"
	"strconv"

	mycsv "github.com/Mirsait/expensia/csv"
	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var outputFile string
var delimiter string

func init() {
	exportCmd.Flags().StringVarP(
		&outputFile,
		OutputFlag,
		"o",
		"expenses.csv",
		"Output File Path")
	exportCmd.Flags().StringVarP(
		&delimiter,
		DelimiterFlag,
		"d",
		",",
		"Delimiter")
	exportCmd.Flags().Bool(
		ForceFlag,
		false,
		"Overwrite output file")
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export data to csv-file",
	Run: func(cmd *cobra.Command, args []string) {
		dataFile := viper.GetString("data")
		if data, err := storage.Load(dataFile); err == nil {
			lst := prepairData(data)
			del := rune(delimiter[0])
			isOverwrite, _ := cmd.Flags().GetBool(ForceFlag)
			err = mycsv.Save(outputFile, lst, del, isOverwrite)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Printf("Data exported to %s\n", outputFile)
		} else {
			fmt.Println("Loading error:", err.Error())
		}
	},
}

func prepairData(items []Expense) [][]string {
	result := make([][]string, len(items)+1)
	result[0] = []string{"DATE", "CATEGORY", "DESCRIPTION", "AMOUNT"}
	for j := 1; j < len(result); j++ {
		result[j] = toList(items[j-1])
	}
	return result
}

func toList(item Expense) []string {
	prettyDate := item.CreatedAt.Format("2006-01-02")
	return []string{
		prettyDate,
		item.Category,
		item.Description,
		strconv.Itoa(item.Amount)}
}
