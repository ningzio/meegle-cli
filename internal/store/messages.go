package store

type TasksRequestedMsg struct {
	ReqID int64
}

type TasksLoadedMsg struct {
	ReqID int64
	Tasks []Task
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

type ApiErrorMsg struct {
	ReqID int64
	Err   error
	Scope string
}
