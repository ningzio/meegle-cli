package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
)

func main() {
	config := app.Config{}
	cmds := meegle.NewCmds(nil, nil)
	application := app.New(config, cmds)
	program := tea.NewProgram(application, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
