package tasks

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
)

func (m *Model) Update(app screen.AppModel, msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case tea.WindowSizeMsg:
		return nil
	}

	return nil
}
