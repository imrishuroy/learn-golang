/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// starttractionCmd represents the starttraction command
var starttractionCmd = &cobra.Command{
	Use:   "starttransaction",
	Short: "Command to start a transaction",
	Long: `This command starts a transaction and returns a transaction id
that is used with stop transaction to id this entire transaction.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := uuid.New()
		fmt.Println("starttransaction called and generated new transaction id", id)
	},
}

func init() {
	rootCmd.AddCommand(starttractionCmd)
}
