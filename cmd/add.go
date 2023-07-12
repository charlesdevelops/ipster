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

// addCmd represents the add command
var (
	// init datastores
	ip   string
	user string
	desc string
	key  string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add an IP address",
		Long: `Add an IP address to the app, along with any other details such as key locations
and descriptions. Use the -d flag to add a description and the -k flag to add a
key location. Example:

ipster add 192.168.0.1 -d "Home router" -u root -k "/home/user/key.pem"`,
		Run: func(cmd *cobra.Command, args []string) {
			// edge cases
			if ip == "" && len(args) != 0 {
				ip = args[0]
			} else {
				fmt.Println("Please specify an IP address. Run 'ipster add --help' for more information.")
				return
			}

			// open db
			db, err := sql.Open("sqlite", "./data.db")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer db.Close()

			// insert logic is dependant on what flags are specified
			switch {
			case ip != "" && user != "" && desc != "" && key != "":
				db.Query("INSERT INTO IPster (ip, user, description, key) VALUES (?, ?, ?, ?)", ip, user, desc, key)
			case ip != "" && user != "" && desc != "":
				db.Query("INSERT INTO IPster (ip, user, description) VALUES (?, ?, ?)", ip, user, desc)
			case ip != "" && user != "" && key != "":
				db.Query("INSERT INTO IPster (ip, user, key) VALUES (?, ?, ?)", ip, user, key)
			case ip != "" && desc != "" && key != "":
				db.Query("INSERT INTO IPster (ip, description, key) VALUES (?, ?, ?)", ip, desc, key)
			case ip != "" && user != "":
				db.Query("INSERT INTO IPster (ip, user) VALUES (?, ?)", ip, user)
			case ip != "" && desc != "":
				db.Query("INSERT INTO IPster (ip, description) VALUES (?, ?)", ip, desc)
			case ip != "" && key != "":
				db.Query("INSERT INTO IPster (ip, key) VALUES (?, ?)", ip, key)
			default:
				db.Query("INSERT INTO IPster (ip) VALUES (?)", ip)
			}

			// // insert logic dependant on what flags are specified
			// switch {
			// case ip != "" && desc != "" && key != "":
			// 	db.Query("INSERT INTO IPster (ip, description, key) VALUES (?, ?, ?)", ip, desc, key)
			// case ip != "" && desc != "":
			// 	db.Query("INSERT INTO IPster (ip, description) VALUES (?, ?)", ip, desc)
			// case ip != "" && key != "":
			// 	db.Query("INSERT INTO IPster (ip, key) VALUES (?, ?)", ip, key)
			// default:
			// 	db.Query("INSERT INTO IPster (ip) VALUES (?)", ip)
			// }
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)

	// stringVarP is a flag that takes a string value, and can be specified with a shorthand
	// ip has an empty shorthand as it is the default item to add
	addCmd.Flags().StringVarP(&ip, "ip", "", "", "Add an IP address, no need to specify the flag by default")
	addCmd.Flags().StringVarP(&user, "user", "u", "", "[Optional] Add the username associated with the address")
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "[Optional] Add a description of the address")
	addCmd.Flags().StringVarP(&key, "key", "k", "", "[Optional] Add the location of the associated SSH key")
}
