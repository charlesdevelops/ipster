/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"database/sql"

	_ "modernc.org/sqlite"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "IPster",
	Short: "A CLI application to store IP addresses in the terminal",
	Long: `IPster is a CLI application that stores IP addresses and optionally, their
associated details, right here in the terminal for quick access. 

To get started, run: ipster add <IP address> -d <description> -k <key location>`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// open db
	db, openErr := sql.Open("sqlite", "./data.db")
	if openErr != nil {
		os.Exit(1)
	}
	defer db.Close()

	// Create table if it doesn't exist
	_, createErr := db.Exec(`CREATE TABLE IF NOT EXISTS IPster (id INTEGER PRIMARY KEY, ip TEXT, key TEXT, description TEXT);`)
	if createErr != nil {
		os.Exit(1)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.IPster.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
