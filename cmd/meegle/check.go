package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"meegle-cli/internal/adapter/auth"
	"meegle-cli/internal/adapter/lark"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if the current login is valid",
	Run: func(cmd *cobra.Command, args []string) {
		tokenStore, err := auth.NewFileTokenStore()
		if err != nil {
			fmt.Printf("Error initializing token store: %v\n", err)
			os.Exit(1)
		}

		client := lark.NewClient(tokenStore)

		if err := client.CheckToken(context.Background()); err != nil {
			fmt.Printf("Token check failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Token is valid (or check completed).")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
