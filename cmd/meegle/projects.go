package main

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"meegle-cli/internal/adapter/auth"
	"meegle-cli/internal/adapter/lark"
	"meegle-cli/internal/ui/model"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Browse and select projects",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Init Auth & Client
		tokenStore, err := auth.NewFileTokenStore()
		if err != nil {
			fmt.Printf("Error initializing token store: %v\n", err)
			os.Exit(1)
		}

		client := lark.NewClient(tokenStore)

		// 2. Fetch Data (Loading state could be in TUI, but for MVP we fetch first)
		fmt.Println("Fetching projects...")
		projects, err := client.ListProjects(context.Background())
		if err != nil {
			fmt.Printf("Failed to list projects: %v\n", err)
			os.Exit(1)
		}

		if len(projects) == 0 {
			fmt.Println("No projects found.")
			return
		}

		// 3. Start TUI
		p := tea.NewProgram(model.NewProjectListModel(projects), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running TUI: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
