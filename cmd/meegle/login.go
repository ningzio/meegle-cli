package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"meegle-cli/internal/adapter/auth"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to Feishu Project",
	Run: func(cmd *cobra.Command, args []string) {
		tokenStore, err := auth.NewFileTokenStore()
		if err != nil {
			fmt.Printf("Error initializing token store: %v\n", err)
			os.Exit(1)
		}

		authenticator := auth.NewLarkAuthenticator(tokenStore)

		fmt.Println("Starting login flow...")
		accessToken, err := authenticator.Login(context.Background())
		if err != nil {
			fmt.Printf("Login failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully logged in! Token (prefix): %s...\n", accessToken[:5])
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
