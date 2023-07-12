/*
Copyright © 2023 CHARLES WATSON
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"database/sql"

	_ "modernc.org/sqlite"
)

var rootCmd = &cobra.Command{
	Use:   "IPster",
	Short: "A CLI application to store IP addresses in the terminal",
	Long: `IPster is a CLI application that stores IP addresses and optionally, their
associated details, right here in the terminal for quick access. 

To get started, run: ipster add <IP address> -d <description> -k <key location>`,
}

func Execute() {

	// open db
	db, openErr := sql.Open("sqlite", "./data.db")
	if openErr != nil {
		os.Exit(1)
	}
	defer db.Close()

	// Create table if it doesn't exist
	_, createErr := db.Exec(`CREATE TABLE IF NOT EXISTS IPster (id INTEGER PRIMARY KEY, ip TEXT, user TEXT, key TEXT, description TEXT);`)
	if createErr != nil {
		os.Exit(1)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
