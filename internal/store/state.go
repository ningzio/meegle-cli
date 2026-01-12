package store

type Task struct {
	ID   string
	Name string
}

type SubTask struct {
	ID     string
	Name   string
	Status string
}

type State struct {
	Tasks               []Task
	TasksByID           map[string]Task
	SubTasksByTaskID    map[string][]SubTask
	SelectedTaskID      string
	SelectedSubTaskID   string
	TasksReqID          int64
	SubTasksReqIDByTask map[string]int64
}

func NewState() State {
	return State{
		TasksByID:           map[string]Task{},
		SubTasksByTaskID:    map[string][]SubTask{},
		SubTasksReqIDByTask: map[string]int64{},
	}
}
