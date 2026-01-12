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
	case SubTasksRequestedMsg:
		s.SubTasksReqIDByTask[m.TaskID] = m.ReqID
	case SubTasksLoadedMsg:
		if s.SubTasksReqIDByTask[m.TaskID] != m.ReqID {
			return s
		}
		s.SubTasksByTaskID[m.TaskID] = m.SubTasks
	case SubTaskCreatedMsg:
		s.SubTasksByTaskID[m.TaskID] = append(s.SubTasksByTaskID[m.TaskID], m.SubTask)
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
