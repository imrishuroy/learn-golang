/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// authorizeCmd represents the authorize command
var authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Check the authroization status of an idTag (OCPP 1.6)",
	Long: `Sends an OCPP 1.6 authorize message
	the central CMS. This will check the authorization status of an idTag`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("idTag")
		if err != nil {
			fmt.Println("Error getting idTag:", err)
		} else {
			fmt.Println("idTag passed in", id)

		}
		fmt.Println("authorize called with params: ", id)
	},
}

func init() {
	rootCmd.AddCommand(authorizeCmd)
	authorizeCmd.Flags().StringP("idTag", "i", "", "Value for ID Tag")

}
