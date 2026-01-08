package tasks

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
)

type Screen struct {
	list    list.Model
	loading bool
	errText string
}

type taskItem struct {
	task meegle.Task
}

func (t taskItem) Title() string       { return t.task.Name }
func (t taskItem) Description() string { return t.task.ID }
func (t taskItem) FilterValue() string { return t.task.Name }

func NewScreen() *Screen {
	delegate := list.NewDefaultDelegate()
	listModel := list.New([]list.Item{}, delegate, 0, 0)
	listModel.Title = "Tasks"
	listModel.SetShowStatusBar(false)
	listModel.SetFilteringEnabled(false)
	listModel.SetShowHelp(false)
	return &Screen{
		list:    listModel,
		loading: true,
	}
}

func (s *Screen) ID() string { return "tasks" }

func (s *Screen) Init() tea.Cmd {
	return func() tea.Msg { return app.TriggerFetchTasksMsg{} }
}
