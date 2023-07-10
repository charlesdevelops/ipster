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

// addCmd represents the add command
var (
	// init datastores
	ip   string
	desc string
	key  string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add an IP address",
		Long: `Add an IP address to the app, along with any other details such as key locations
and descriptions. Use the -d flag to add a description and the -k flag to add a
key location. Example:

ipster add 192.168.0.1 -d "Home router" -k "/home/user/key.pem"`,
		Run: func(cmd *cobra.Command, args []string) {
			if ip == "" && len(args) != 0 {
				ip = args[0]
			} else {
				fmt.Println("Please specify an IP address. Run 'ipster add --help' for more information.")
				return
			}

			db, err := sql.Open("sqlite", "./data.db")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer db.Close()

			switch {
			case ip != "" && desc != "" && key != "":
				db.Query("INSERT INTO IPster (ip, description, key) VALUES (?, ?, ?)", ip, desc, key)
			case ip != "" && desc != "":
				db.Query("INSERT INTO IPster (ip, description) VALUES (?, ?)", ip, desc)
			case ip != "" && key != "":
				db.Query("INSERT INTO IPster (ip, key) VALUES (?, ?)", ip, key)
			default:
				db.Query("INSERT INTO IPster (ip) VALUES (?)", ip)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&ip, "ip", "", "", "Add an IP address, no need to specify the flag by default")
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "[Optional] Add a description of the address")
	addCmd.Flags().StringVarP(&key, "key", "k", "", "[Optional] Add the location of the associated SSH key")
}
