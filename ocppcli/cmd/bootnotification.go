/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bootnotificationCmd represents the bootnotification command
var bootnotificationCmd = &cobra.Command{
	Use:   "bootnotification",
	Short: "Charger bootnotification (OCPP 1.6)",
	Long: `Sent when the charger boots up, contains 
	info such as FW version, vendor/model, 
	serial number, etc.`,
	Args:    cobra.ExactArgs(1),
	Example: "occpcli bootnotification up -s 12323 -m abcd",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := cmd.Flags().GetString("serialNumber")
		if err != nil {
			fmt.Println("Error getting serialNumber:", err)
		}

		m, err := cmd.Flags().GetString("make")
		if err != nil {
			fmt.Println("Error getting make:", err)

		}
		fmt.Println("bootnotification called with arg: ", args[0], "and with params", s, "/", m)
	},
}

func init() {
	rootCmd.AddCommand(bootnotificationCmd)

	bootnotificationCmd.Flags().StringP("serialNumber", "s", "", "Serial Number of Unit")
	bootnotificationCmd.Flags().StringP("make", "m", "", "Make of Unit")
	bootnotificationCmd.MarkFlagRequired("serialNumber")
	bootnotificationCmd.MarkFlagRequired("make")

}
