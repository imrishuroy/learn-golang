/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// metervalueCmd represents the metervalue command
var metervalueCmd = &cobra.Command{
	Use:   "metervalue",
	Short: "Meter Value command",
	Long: `A meter value that is sent as part of a transaction
   containing one key which is the id or type and a second flag for value that is
   is the actual measurement or value to be transmitted.`,
	Run: func(cmd *cobra.Command, args []string) {
		k, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println("Error getting key:", err)
			panic(err)
		}
		v, err := cmd.Flags().GetString("value")
		if err != nil {
			fmt.Println("Error getting value:", err)
			panic(err)
		}
		t, err := cmd.Flags().GetString("transaction")
		if err != nil {
			fmt.Println("Error getting transaction:", err)
			panic(err)
		}
		fmt.Println("metervalue for transaction", t, "called with params", k, "/", v)
	},
}

func init() {
	rootCmd.AddCommand(metervalueCmd)

	metervalueCmd.Flags().StringP("key", "k", "", "Key for this meter value. e.g. KwH")
	metervalueCmd.Flags().StringP("value", "v", "", "Value for this meter value. e.g. 380")
	metervalueCmd.Flags().StringP("transaction", "t", "", "Transaction id for this meter value")
	metervalueCmd.MarkFlagRequired("transaction")
	metervalueCmd.MarkFlagsRequiredTogether("key", "value") // require both key and value

}
