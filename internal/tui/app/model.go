package app

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ningzio/meegle-cli/internal/model"
	"github.com/ningzio/meegle-cli/internal/service"
)

// Msg types
type tasksLoadedMsg []model.Task
type errMsg error

// Model represents the main application state.
type Model struct {
	service *service.TaskService
	tasks   []model.Task
	err     error
}

// NewModel creates a new main Model.
func NewModel(svc *service.TaskService) Model {
	return Model{service: svc}
}

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	return m.fetchTasks
}

// fetchTasks is a command that fetches tasks from the service.
func (m Model) fetchTasks() tea.Msg {
	tasks, err := m.service.ListTasks(context.Background())
	if err != nil {
		return errMsg(err)
	}
	return tasksLoadedMsg(tasks)
}

// Update handles messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tasksLoadedMsg:
		m.tasks = msg
		m.err = nil
	case errMsg:
		m.err = msg
	}
	return m, nil
}

// View renders the UI.
func (m Model) View() string {
	s := "Meegle CLI\n\n"
	if m.err != nil {
		s += "Error: " + m.err.Error() + "\n"
	} else {
		if len(m.tasks) == 0 {
			s += "Loading...\n"
		} else {
			for _, t := range m.tasks {
				s += "- " + t.Title + " [" + t.Status + "]\n"
			}
		}
	}
	s += "\nPress 'q' to quit.\n"
	return s
}
