package editor

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
)

type Mode int

const (
	ModeTask Mode = iota
	ModeSubTask
)

type Model struct {
	Input textinput.Model
	Mode  Mode
}

func NewTask() *Model {
	return newEditor(ModeTask, "New task name")
}

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
