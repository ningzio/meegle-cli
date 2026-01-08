package meegle

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/store"
)

func projectKey() string {
	key := os.Getenv("MEEGLE_PROJECT_KEY")
	if key == "" {
		return "demo"
	}
	return key
}

func FetchTasksCmd(client Client, reqID int) tea.Cmd {
	return func() tea.Msg {
		tasks, err := client.FetchTasks(projectKey())
		return store.TasksFetchedMsg{ReqID: reqID, Tasks: tasks, Err: err}
	}
}

func CreateTaskCmd(client Client, reqID int, name string) tea.Cmd {
	return func() tea.Msg {
		task, err := client.CreateTask(projectKey(), name)
		return store.TaskCreatedMsg{ReqID: reqID, Task: task, Err: err}
	}
}

func CreateSubTaskCmd(client Client, reqID int, taskID, name string) tea.Cmd {
	return func() tea.Msg {
		sub, err := client.CreateSubTask(projectKey(), taskID, name)
		return store.SubTaskCreatedMsg{ReqID: reqID, TaskID: taskID, SubTask: sub, Err: err}
	}
}

func ToggleSubTaskDoneCmd(client Client, reqID int, taskID, subTaskID string, done bool) tea.Cmd {
	return func() tea.Msg {
		sub, err := client.ToggleSubTaskDone(projectKey(), taskID, subTaskID, done)
		return store.SubTaskToggledMsg{ReqID: reqID, TaskID: taskID, SubTask: sub, Err: err}
	}
}
