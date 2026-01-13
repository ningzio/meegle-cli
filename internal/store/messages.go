package store

type TasksRequestedMsg struct {
	ReqID int64
}

type TasksLoadedMsg struct {
	ReqID int64
	Tasks []Task
}

type TaskCreatedMsg struct {
	Task Task
}

type TasksFailedMsg struct {
	ReqID int64
	Err   error
}

type SubTasksRequestedMsg struct {
	ReqID  int64
	TaskID string
}

type SubTasksLoadedMsg struct {
	ReqID    int64
	TaskID   string
	SubTasks []SubTask
}

type SubTaskCreatedMsg struct {
	TaskID  string
	SubTask SubTask
}

type SubTaskCompletedMsg struct {
	TaskID    string
	SubTaskID string
}

type SubTaskRolledBackMsg struct {
	TaskID    string
	SubTaskID string
}

type TaskSelectedMsg struct {
	TaskID string
}

type SubTaskSelectedMsg struct {
	TaskID    string
	SubTaskID string
}

type APIErrorMsg struct {
	ReqID int64
	Err   error
	Scope string
}
