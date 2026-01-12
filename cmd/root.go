package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expensia",
	Short: "Expense tracker",
	Long: `This is a simple command-line application designed to help users manage their
personal finances. The application allows users to add, update, delete, and view
expenses.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Author:", viper.GetString("author"))
		fmt.Println("License:", viper.GetString("license"))
		fmt.Println("Data:", viper.GetString("data"))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(
		&dataFile,
		DataFlag,
		"",
		"data file (default is $HOME/expensia/data.yaml)")
	viper.BindPFlag("data", rootCmd.Flags().Lookup("data"))

	// set if none
	viper.SetDefault("author", "Aucharenka Mikhail <mavcharenko@gmail.com>")
	viper.SetDefault("license", "MIT")
	viper.SetDefault("data", getDefaultPath())
}

func getDefaultPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, "expensia", "data.json")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		configPath := filepath.Join(home, "expensia")
		_ = os.MkdirAll(configPath, 0755)

		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
		cfgFile = filepath.Join(configPath, ".config.yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config not found, creating:", cfgFile)
			if err := viper.SafeWriteConfigAs(cfgFile); err != nil {
				cobra.CheckErr(err)
			}
		} else {
			cobra.CheckErr(err)
		}
	} else {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
		if !viper.IsSet("data") {
			viper.Set("data", getDefaultPath())
		}

		if err := viper.WriteConfigAs(cfgFile); err != nil {
			cobra.CheckErr(err)
		}
	}
}
