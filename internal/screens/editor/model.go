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

// Init prepares the editor for first render.
func (m *Model) Init(app screen.AppModel) tea.Cmd {
	return textinput.Blink
}

// OnFocus refreshes the editor when the screen gains focus.
func (m *Model) OnFocus(app screen.AppModel) tea.Cmd {
	return textinput.Blink
}

// OnBlur handles editor teardown when the screen loses focus.
func (m *Model) OnBlur(app screen.AppModel) {}
