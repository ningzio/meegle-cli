package store

import tea "github.com/charmbracelet/bubbletea"

func Reduce(s State, msg tea.Msg) State {
	switch m := msg.(type) {
	case TasksRequestedMsg:
		s.TasksReqID = m.ReqID
	case TasksLoadedMsg:
		if m.ReqID != s.TasksReqID {
			return s
		}
		s.Tasks = m.Tasks
		s.TasksByID = indexTasks(m.Tasks)
	case TaskCreatedMsg:
		s.Tasks = append(s.Tasks, m.Task)
		s.TasksByID[m.Task.ID] = m.Task
	case SubTasksRequestedMsg:
		s.SubTasksReqIDByTask[m.TaskID] = m.ReqID
	case SubTasksLoadedMsg:
		if s.SubTasksReqIDByTask[m.TaskID] != m.ReqID {
			return s
		}
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
