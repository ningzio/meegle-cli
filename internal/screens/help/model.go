package help

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/ui/components/help"
)

// Model represents the help screen state.
type Model struct {
	overlay help.Model
}

// New constructs a help screen with default shortcuts.
func New() *Model {
	return &Model{overlay: help.New(defaultLines())}
}

// Init has no startup behavior for the help screen.
func (m *Model) Init(_ screen.AppModel) tea.Cmd {
	return nil
}

// OnFocus shows the help overlay.
func (m *Model) OnFocus(_ screen.AppModel) tea.Cmd {
	m.overlay.Visible = true
	return nil
}

// OnBlur hides the help overlay.
func (m *Model) OnBlur(_ screen.AppModel) {
	m.overlay.Visible = false
}

func defaultLines() []string {
	return []string{
		"Global",
		"  ?: help",
		"  esc: back",
		"  q: quit",
		"",
		"Tasks",
		"  n: new task",
		"  enter: open task",
		"",
		"Task detail",
		"  a: add subtask",
		"  c: complete subtask",
		"  r: rollback subtask",
	}
}
