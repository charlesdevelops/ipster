/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
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
		db, err := sql.Open("sqlite", "./data.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		rows, err := db.Query("SELECT * FROM IPster")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var id int
		var ip, key string
		for rows.Next() {
			err := rows.Scan(&id, &ip, &key)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%d: %s, %s\n", id, ip, key)
		}

		if err := rows.Err(); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
