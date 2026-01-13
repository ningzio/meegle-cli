package store

import tea "github.com/charmbracelet/bubbletea"

// Reduce applies a message to the state for the update loop.
// It is not concurrency-safe and should be called from the Bubble Tea update goroutine.
func Reduce(s State, msg tea.Msg) State {
	switch m := msg.(type) {
	case TasksRequestedMsg:
		s.TasksReqID = m.ReqID
	case TasksLoadedMsg:
		if s.TasksReqID != 0 && m.ReqID != s.TasksReqID {
			return s
		}
		s.TasksReqID = m.ReqID
		s.Tasks = m.Tasks
		s.TasksByID = indexTasks(m.Tasks)
	case TaskCreatedMsg:
		s.Tasks = append(s.Tasks, m.Task)
		s.TasksByID[m.Task.ID] = m.Task
	case SubTasksRequestedMsg:
		s.SubTasksReqIDByTask[m.TaskID] = m.ReqID
	case SubTasksLoadedMsg:
		currentReqID := s.SubTasksReqIDByTask[m.TaskID]
		if currentReqID != 0 && currentReqID != m.ReqID {
			return s
		}
		s.SubTasksReqIDByTask[m.TaskID] = m.ReqID
		s.SubTasksByTaskID[m.TaskID] = m.SubTasks
	case SubTaskCreatedMsg:
		s.SubTasksByTaskID[m.TaskID] = append(s.SubTasksByTaskID[m.TaskID], m.SubTask)
	case SubTaskCompletedMsg:
		s.SubTasksByTaskID[m.TaskID] = updateSubTaskStatus(s.SubTasksByTaskID[m.TaskID], m.SubTaskID, "completed")
	case SubTaskRolledBackMsg:
		s.SubTasksByTaskID[m.TaskID] = updateSubTaskStatus(s.SubTasksByTaskID[m.TaskID], m.SubTaskID, "open")
	case TaskSelectedMsg:
		s.SelectedTaskID = m.TaskID
		s.SelectedSubTaskID = ""
	case SubTaskSelectedMsg:
		s.SelectedTaskID = m.TaskID
		s.SelectedSubTaskID = m.SubTaskID
	}

	return s
}

func indexTasks(tasks []Task) map[string]Task {
	indexed := make(map[string]Task, len(tasks))
	for _, task := range tasks {
		indexed[task.ID] = task
	}
	return indexed
}

func updateSubTaskStatus(subTasks []SubTask, subTaskID, status string) []SubTask {
	updated := make([]SubTask, 0, len(subTasks))
	for _, subTask := range subTasks {
		if subTask.ID == subTaskID {
			subTask.Status = status
		}
		updated = append(updated, subTask)
	}
	return updated
}
