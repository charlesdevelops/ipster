/*
Copyright © 2023 NAME HERE CHARLES WATSON
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"database/sql"

	_ "modernc.org/sqlite"
)

// rmCmd represents the rm command
var (
	// true if removing all data
	all   bool
	rmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove entries from the application",
		Long: `rm or Remove will remove an entry from the application, or alternatively
remove all entries by using "rm -a". Example:

ipster rm 1`,
		Run: func(cmd *cobra.Command, args []string) {
			// open db
			db, err := sql.Open("sqlite", "./data.db")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer db.Close()

			if all {
				_, err := db.Query("DROP TABLE IF EXISTS IPster")
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("All entries removed.")
				return
			} else {
				db.Query("DELETE FROM IPster WHERE id = ?", args[0])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(rmCmd)
	rmCmd.Flags().BoolVarP(&all, "all", "a", false, "Removes all entries from the application")
}
