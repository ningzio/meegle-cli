package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
	"meegle-cli/internal/screens/tasks"
	"meegle-cli/internal/store"
)

func main() {
	client := meegle.NewClientFromEnv()
	state := store.NewState()
	initial := tasks.NewScreen()
	router := app.NewRouter(initial)
	model := app.NewModel(router, client, state)

	program := tea.NewProgram(model, tea.WithAltScreen())
	if err := program.Start(); err != nil {
		log.Fatal(err)
	}
}
