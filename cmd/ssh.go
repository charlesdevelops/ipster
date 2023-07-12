/*
Copyright © 2023 CHARLES WATSON
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/atotto/clipboard"

	"database/sql"

	_ "modernc.org/sqlite"
)

// sshCmd represents the ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Generates an ssh command to a saved IP address",
	Long: `Generates an ssh command to a saved IP address, including any key location
if one has been specified. The command is printed to the terminal,
as well as being saved to your clipboard. Usage:

ipster ssh [1]`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			fmt.Println("Include a valid IP address ID number.")
			return
		}

		// open db
		db, err := sql.Open("sqlite", "./data.db")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()

		var ip, key, user sql.NullString
		var outputStr string

		queryData, err := db.Query("SELECT ip, key, user FROM IPster WHERE id = ?", args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer queryData.Close()
		queryData.Next()
		queryData.Scan(&ip, &key, &user)

		switch {
		case key.Valid && ip.Valid && user.Valid:
			outputStr = fmt.Sprintf("ssh -i %s %s@%s", key.String, user.String, ip.String)
		case ip.Valid && user.Valid:
			outputStr = fmt.Sprintf("ssh %s@%s", user.String, ip.String)
		case ip.Valid:
			outputStr = fmt.Sprintf("ssh %s", ip.String)
		default:
			fmt.Println("No IP address found.")
		}

		err = clipboard.WriteAll(string(outputStr))
		if err != nil {
			fmt.Println("Clipboard error:", err)
			return
		} else {
			fmt.Println("Copied to clipboard:", outputStr)
		}
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
