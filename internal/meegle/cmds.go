package meegle

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/store"
)

// Cmds provides Bubble Tea command factories for Meegle operations.
type Cmds struct {
	client *Client
	auth   *AuthManager
}

// NewCmds constructs a command factory for Meegle requests.
func NewCmds(client *Client, auth *AuthManager) *Cmds {
	return &Cmds{client: client, auth: auth}
}

func (c *Cmds) FetchTasks(_ string, reqID int64) tea.Cmd {
	return func() tea.Msg {
		tasks := []store.Task{
			{ID: "task-1", Name: "Design Milestone 1"},
			{ID: "task-2", Name: "Build TUI Skeleton"},
			{ID: "task-3", Name: "Ship MVP Flow"},
			{ID: "task-4", Name: "User Onboarding Copy"},
			{ID: "task-5", Name: "Accessibility Sweep"},
			{ID: "task-6", Name: "Release Notes Draft"},
		}
		return store.TasksLoadedMsg{ReqID: reqID, Tasks: tasks}
	}
}

func (c *Cmds) FetchSubTasks(_, taskID string, reqID int64) tea.Cmd {
	return func() tea.Msg {
		subTasks := []store.SubTask{
			{ID: fmt.Sprintf("%s-sub-1", taskID), Name: "Draft plan", Status: "open"},
			{ID: fmt.Sprintf("%s-sub-2", taskID), Name: "Review details", Status: "open"},
			{ID: fmt.Sprintf("%s-sub-3", taskID), Name: "Execute work", Status: "completed"},
			{ID: fmt.Sprintf("%s-sub-4", taskID), Name: "QA pass", Status: "open"},
			{ID: fmt.Sprintf("%s-sub-5", taskID), Name: "Ship checklist", Status: "completed"},
		}
		return store.SubTasksLoadedMsg{ReqID: reqID, TaskID: taskID, SubTasks: subTasks}
	}
}

func (c *Cmds) CreateTask(_, name string) tea.Cmd {
	return func() tea.Msg {
		return store.TaskCreatedMsg{Task: store.Task{ID: fmt.Sprintf("task-%d", time.Now().UnixNano()), Name: name}}
	}
}

func (c *Cmds) CreateSubTask(_, taskID, name string) tea.Cmd {
	return func() tea.Msg {
		subTask := store.SubTask{
			ID:     fmt.Sprintf("%s-sub-%d", taskID, time.Now().UnixNano()),
			Name:   name,
			Status: "open",
		}
		return store.SubTaskCreatedMsg{TaskID: taskID, SubTask: subTask}
	}
}

func (c *Cmds) CompleteSubTask(_, taskID, subTaskID string) tea.Cmd {
	return func() tea.Msg {
		return store.SubTaskCompletedMsg{TaskID: taskID, SubTaskID: subTaskID}
	}
}

func (c *Cmds) RollbackSubTask(_, taskID, subTaskID string) tea.Cmd {
	return func() tea.Msg {
		return store.SubTaskRolledBackMsg{TaskID: taskID, SubTaskID: subTaskID}
	}
}
