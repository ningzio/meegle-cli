package store

import "meegle-cli/internal/meegle"

type TasksFetchedMsg struct {
	ReqID int
	Tasks []meegle.Task
	Err   error
}

type TaskCreatedMsg struct {
	ReqID int
	Task  meegle.Task
	Err   error
}

type SubTaskCreatedMsg struct {
	ReqID   int
	TaskID  string
	SubTask meegle.SubTask
	Err     error
}

type SubTaskToggledMsg struct {
	ReqID   int
	TaskID  string
	SubTask meegle.SubTask
	Err     error
}

func Reduce(state State, msg interface{}) State {
	switch typed := msg.(type) {
	case TasksFetchedMsg:
		if typed.Err != nil || !IsLatest(state, ReqFetchTasks, typed.ReqID) {
			return state
		}
		state.Tasks = typed.Tasks
		if len(state.Tasks) > 0 {
			state.SelectedTaskID = state.Tasks[0].ID
		}
	case TaskCreatedMsg:
		if typed.Err != nil || !IsLatest(state, ReqCreateTask, typed.ReqID) {
			return state
		}
		state.Tasks = append([]meegle.Task{typed.Task}, state.Tasks...)
		state.SelectedTaskID = typed.Task.ID
	case SubTaskCreatedMsg:
		if typed.Err != nil || !IsLatest(state, ReqCreateSubTask, typed.ReqID) {
			return state
		}
		state.SubTasks[typed.TaskID] = append([]meegle.SubTask{typed.SubTask}, state.SubTasks[typed.TaskID]...)
	case SubTaskToggledMsg:
		if typed.Err != nil || !IsLatest(state, ReqToggleSubTask, typed.ReqID) {
			return state
		}
		items := state.SubTasks[typed.TaskID]
		for i, item := range items {
			if item.ID == typed.SubTask.ID {
				items[i] = typed.SubTask
				break
			}
		}
		state.SubTasks[typed.TaskID] = items
	}
	return state
}
