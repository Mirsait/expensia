package cmd

import (
	"fmt"

	"github.com/Mirsait/expensia/storage"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show expenses list",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		datafile := viper.GetString("data")
		if data, err := storage.Load(datafile); err == nil {
			fmt.Printf("%4s %16s %12s %6s\n", "ID", "Date", "Description", "Amount")
			for _, e := range data {
				prettyDate := e.CreatedAt.Format("02-01-2006 15:04")
				fmt.Printf("%04d %16s %12s $%d\n",
					e.Id, prettyDate, e.Description, e.Amount)
			}
		}
	},
}
