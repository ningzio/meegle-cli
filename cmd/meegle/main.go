package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"meegle-cli/internal/adapter/config"
)

var rootCmd = &cobra.Command{
	Use:   "meegle",
	Short: "Meegle CLI - A TUI for Feishu Project",
	Long: `Meegle CLI is a terminal-based interactive tool for managing Feishu Project tasks.
It focuses on "Context-Aware" interactions and efficiency.`,
}

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing config: %v\n", err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
