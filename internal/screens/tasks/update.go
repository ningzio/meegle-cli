package tasks

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
)

func (m *Model) Update(app *app.App, msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case tea.WindowSizeMsg:
		return nil
	}

	return nil
}
