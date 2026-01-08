package store

import "meegle-cli/internal/meegle"

func Reduce(state State, msg interface{}) State {
	switch typed := msg.(type) {
	case meegle.TasksFetchedMsg:
		if typed.Err != nil || !IsLatest(state, ReqFetchTasks, typed.ReqID) {
			return state
		}
		state.Tasks = typed.Tasks
		if len(state.Tasks) > 0 {
			state.SelectedTaskID = state.Tasks[0].ID
		}
	case meegle.TaskCreatedMsg:
		if typed.Err != nil || !IsLatest(state, ReqCreateTask, typed.ReqID) {
			return state
		}
		state.Tasks = append([]meegle.Task{typed.Task}, state.Tasks...)
		state.SelectedTaskID = typed.Task.ID
	case meegle.SubTaskCreatedMsg:
		if typed.Err != nil || !IsLatest(state, ReqCreateSubTask, typed.ReqID) {
			return state
		}
		state.SubTasks[typed.TaskID] = append([]meegle.SubTask{typed.SubTask}, state.SubTasks[typed.TaskID]...)
	case meegle.SubTaskToggledMsg:
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
