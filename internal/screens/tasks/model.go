package tasks

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
)

type Model struct {
	Loading bool
}

func New() *Model {
	return &Model{Loading: true}
}

func (m *Model) Init(app *app.App) tea.Cmd {
	return nil
}

func (m *Model) OnFocus(app *app.App) tea.Cmd {
	return nil
}

func (m *Model) OnBlur(app *app.App) {}
