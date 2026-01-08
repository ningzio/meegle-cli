package store

import "meegle-cli/internal/meegle"

func SelectedTask(state State) *meegle.Task {
	if state.SelectedTaskID == "" {
		return nil
	}
	for _, task := range state.Tasks {
		if task.ID == state.SelectedTaskID {
			copy := task
			return &copy
		}
	}
	return nil
}

func SubTasksFor(state State, taskID string) []meegle.SubTask {
	items := state.SubTasks[taskID]
	return append([]meegle.SubTask{}, items...)
}
