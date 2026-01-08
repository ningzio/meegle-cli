package detail

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
)

type Screen struct {
	taskID string
	list   list.Model
}

type subTaskItem struct {
	subTask meegle.SubTask
}

func (s subTaskItem) Title() string {
	status := ""
	if s.subTask.Done {
		status = "✓ "
	}
	return status + s.subTask.Name
}

func (s subTaskItem) Description() string { return s.subTask.ID }
func (s subTaskItem) FilterValue() string { return s.subTask.Name }

func NewScreen(taskID string) *Screen {
	delegate := list.NewDefaultDelegate()
	listModel := list.New([]list.Item{}, delegate, 0, 0)
	listModel.Title = "Subtasks"
	listModel.SetShowStatusBar(false)
	listModel.SetFilteringEnabled(false)
	listModel.SetShowHelp(false)
	return &Screen{taskID: taskID, list: listModel}
}

func (s *Screen) ID() string { return "detail" }

func (s *Screen) Init() tea.Cmd {
	return nil
}
