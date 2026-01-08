package editor

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
)

type Mode int

const (
	ModeTask Mode = iota
	ModeSubTask
)

type Screen struct {
	mode       Mode
	taskID     string
	input      textinput.Model
	submitting bool
}

func NewTaskScreen() *Screen {
	input := textinput.New()
	input.Placeholder = "Task name"
	input.Focus()
	input.CharLimit = 120
	input.Width = 30
	return &Screen{mode: ModeTask, input: input}
}

func NewSubTaskScreen(taskID string) *Screen {
	input := textinput.New()
	input.Placeholder = "Subtask name"
	input.Focus()
	input.CharLimit = 120
	input.Width = 30
	return &Screen{mode: ModeSubTask, taskID: taskID, input: input}
}

func (s *Screen) ID() string { return "editor" }

func (s *Screen) Init() tea.Cmd { return textinput.Blink }
