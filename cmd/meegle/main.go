package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ningzio/meegle-cli/internal/adapter/api"
	"github.com/ningzio/meegle-cli/internal/service"
	"github.com/ningzio/meegle-cli/internal/tui/app"
)

func main() {
	client := api.NewMockClient()
	svc := service.NewTaskService(client)
	m := app.NewModel(svc)

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error starting program: %v\n", err)
		os.Exit(1)
	}
}
