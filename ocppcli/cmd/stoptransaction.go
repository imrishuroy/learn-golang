/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stoptransactionCmd represents the stoptransaction command
var stoptransactionCmd = &cobra.Command{
	Use:   "stoptransaction",
	Short: "Stop transaction command",
	Long: `Called to signal a transaction containing zero or more
   meter values has ended.`,
	Run: func(cmd *cobra.Command, args []string) {
		t, err := cmd.Flags().GetString("transaction")
		if err != nil {
			fmt.Println("Error getting transaction:", err)
			panic(err)
		}
		fmt.Println("stoptransaction for transaction: ", t, "called")
	},
}

func init() {
	rootCmd.AddCommand(stoptransactionCmd)

	rootCmd.AddCommand(stoptransactionCmd)

	stoptransactionCmd.Flags().StringP("transaction", "t", "", "Transaction Id to stop")
	stoptransactionCmd.MarkFlagRequired("transaction")

}
