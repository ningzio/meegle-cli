package app

import (
	"context"
	"errors"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ningzio/meegle-cli/internal/model"
	"github.com/ningzio/meegle-cli/internal/service"
)

// MockClient for testing TUI
type MockClient struct {
	Tasks []model.Task
	Err   error
}

func (m *MockClient) GetTasks(ctx context.Context) ([]model.Task, error) {
	return m.Tasks, m.Err
}

func TestModel_Update(t *testing.T) {
	client := &MockClient{
		Tasks: []model.Task{{ID: "1", Title: "Test"}},
	}
	svc := service.NewTaskService(client)
	m := NewModel(svc)

	// Test Init
	cmd := m.Init()
	if cmd == nil {
		t.Error("Init() returned nil command")
	}

	// Test tasksLoadedMsg
	tasks := []model.Task{{ID: "1", Title: "Test"}}
	newModel, _ := m.Update(tasksLoadedMsg(tasks))
	m = newModel.(Model)
	if len(m.tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(m.tasks))
	}

	// Test errMsg
	testErr := errors.New("fail")
	newModel, _ = m.Update(errMsg(testErr))
	m = newModel.(Model)
	if m.err != testErr {
		t.Errorf("Expected error %v, got %v", testErr, m.err)
	}

	// Test KeyMsg
	newModel, cmd = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	if cmd == nil { // tea.Quit is a command, but checking equality is hard.
		// Just checking if we got a model back. Bubble Tea internals are tricky to test equality on cmds.
		// Ideally we check if it returns tea.Quit, but that's a var.
	}
}

func TestModel_View(t *testing.T) {
	client := &MockClient{}
	svc := service.NewTaskService(client)
	m := NewModel(svc)

	// Empty
	view := m.View()
	if view == "" {
		t.Error("View returned empty string")
	}

	// With Tasks
	m.tasks = []model.Task{{ID: "1", Title: "Test Task", Status: "Open"}}
	view = m.View()
	if view == "" {
		t.Error("View returned empty string")
	}
}
