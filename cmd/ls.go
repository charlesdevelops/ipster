/*
Copyright © 2023 CHARLES WATSON
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"database/sql"

	_ "modernc.org/sqlite"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all saved addresses",
	Long: `This command lists out all the IP addresses you have saved,
including any other details such as key locations and descriptions.`,

	Run: func(cmd *cobra.Command, args []string) {

		// open db
		db, err := sql.Open("sqlite", "./data.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		// select 1 row at a time
		rows, err := db.Query("SELECT id, ip, key, description FROM IPster")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		fmt.Println("ID | IP | Description | Key Location")
		fmt.Println("----------------------------")

		// declare variables to store data from db
		// nullString is a special type that can store empty values, avoiding panics
		var id int
		var ip, key, description sql.NullString

		// iterate through rows 1 at a time
		for rows.Next() {
			err := rows.Scan(&id, &ip, &key, &description)
			if err != nil {
				fmt.Println(err)
				return
			}

			// print row logic
			switch {
			case key.Valid && description.Valid:
				fmt.Printf("%d: %s | %s | %s\n", id, ip.String, description.String, key.String)
			case key.Valid:
				fmt.Printf("%d: %s | - | %s\n", id, ip.String, key.String)
			case description.Valid:
				fmt.Printf("%d: %s | %s | - \n", id, ip.String, description.String)
			default:
				fmt.Printf("%d: %s\n", id, ip.String)
			}

			if err := rows.Err(); err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
