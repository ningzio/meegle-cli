package meegle

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/store"
)

type Cmds struct {
	client *Client
	auth   *AuthManager
}

func NewCmds(client *Client, auth *AuthManager) *Cmds {
	return &Cmds{client: client, auth: auth}
}

func (c *Cmds) FetchTasks(projectKey string, reqID int64) tea.Cmd {
	return func() tea.Msg {
		return store.ApiErrorMsg{ReqID: reqID, Err: errors.New("fetch tasks not implemented"), Scope: "tasks"}
	}
}

func (c *Cmds) FetchSubTasks(projectKey, taskID string, reqID int64) tea.Cmd {
	return func() tea.Msg {
		return store.ApiErrorMsg{ReqID: reqID, Err: errors.New("fetch sub tasks not implemented"), Scope: "subtasks"}
	}
}

func (c *Cmds) CreateTask(projectKey, name string) tea.Cmd {
	return func() tea.Msg {
		return store.ApiErrorMsg{Err: errors.New("create task not implemented"), Scope: "tasks"}
	}
}

func (c *Cmds) CreateSubTask(projectKey, taskID, name string) tea.Cmd {
	return func() tea.Msg {
		return store.ApiErrorMsg{Err: errors.New("create sub task not implemented"), Scope: "subtasks"}
	}
}

func (c *Cmds) CompleteSubTask(projectKey, taskID, subTaskID string) tea.Cmd {
	return func() tea.Msg {
		return store.ApiErrorMsg{Err: errors.New("complete sub task not implemented"), Scope: "subtasks"}
	}
}

func (c *Cmds) RollbackSubTask(projectKey, taskID, subTaskID string) tea.Cmd {
	return func() tea.Msg {
		return store.ApiErrorMsg{Err: errors.New("rollback sub task not implemented"), Scope: "subtasks"}
	}
}
