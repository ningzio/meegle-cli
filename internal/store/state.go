package store

// Task represents a task in the store layer.
type Task struct {
	ID   string
	Name string
}

// SubTask represents a subtask in the store layer.
type SubTask struct {
	ID     string
	Name   string
	Status string
}

// State holds the current application data snapshot.
type State struct {
	Tasks               []Task
	TasksByID           map[string]Task
	SubTasksByTaskID    map[string][]SubTask
	SelectedTaskID      string
	SelectedSubTaskID   string
	TasksReqID          int64
	SubTasksReqIDByTask map[string]int64
}

// NewState returns an initialized store state with ready-to-use maps.
func NewState() State {
	return State{
		TasksByID:           map[string]Task{},
		SubTasksByTaskID:    map[string][]SubTask{},
		SubTasksReqIDByTask: map[string]int64{},
	}
}
