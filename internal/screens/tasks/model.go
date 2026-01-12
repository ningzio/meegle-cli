package tasks

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
)

type Model struct {
	Loading bool
}

func New() *Model {
	return &Model{Loading: true}
}

func (m *Model) Init(app screen.AppModel) tea.Cmd {
	return nil
}

func (m *Model) OnFocus(app screen.AppModel) tea.Cmd {
	return nil
}

func (m *Model) OnBlur(app screen.AppModel) {}
