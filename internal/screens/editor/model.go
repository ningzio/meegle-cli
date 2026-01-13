package editor

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
)

// Mode describes which entity the editor is creating.
type Mode int

const (
	// ModeTask configures the editor for creating tasks.
	ModeTask Mode = iota
	// ModeSubTask configures the editor for creating subtasks.
	ModeSubTask
)

// Model represents the editor screen state.
type Model struct {
	Input textinput.Model
	Mode  Mode
}

// NewTask returns an editor model configured for task creation.
func NewTask() *Model {
	return newEditor(ModeTask, "New task name")
}

// NewSubTask returns an editor model configured for subtask creation.
func NewSubTask() *Model {
	return newEditor(ModeSubTask, "New subtask name")
}

func newEditor(mode Mode, placeholder string) *Model {
	input := textinput.New()
	input.Placeholder = placeholder
	input.Focus()
	return &Model{Input: input, Mode: mode}
}

func (m *Model) Init(_ screen.AppModel) tea.Cmd {
	return textinput.Blink
}

func (m *Model) OnFocus(_ screen.AppModel) tea.Cmd {
	return textinput.Blink
}

func (m *Model) OnBlur(_ screen.AppModel) {}
