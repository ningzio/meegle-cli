package store

// TasksRequestedMsg signals that tasks should be loaded for a request ID.
type TasksRequestedMsg struct {
	ReqID int64
}

// TasksLoadedMsg delivers the tasks associated with a request ID.
type TasksLoadedMsg struct {
	ReqID int64
	Tasks []Task
}

// TaskCreatedMsg announces that a task was created.
type TaskCreatedMsg struct {
	Task Task
}

// TasksFailedMsg reports a failure while loading tasks for a request ID.
type TasksFailedMsg struct {
	ReqID int64
	Err   error
}

// SubTasksRequestedMsg signals that subtasks should be loaded for a task.
type SubTasksRequestedMsg struct {
	ReqID  int64
	TaskID string
}

// SubTasksLoadedMsg delivers subtasks for a task and request ID.
type SubTasksLoadedMsg struct {
	ReqID    int64
	TaskID   string
	SubTasks []SubTask
}

// SubTaskCreatedMsg announces that a subtask was created for a task.
type SubTaskCreatedMsg struct {
	TaskID  string
	SubTask SubTask
}

// SubTaskCompletedMsg marks a subtask as completed for a task.
type SubTaskCompletedMsg struct {
	TaskID    string
	SubTaskID string
}

// SubTaskRolledBackMsg marks a subtask as rolled back for a task.
type SubTaskRolledBackMsg struct {
	TaskID    string
	SubTaskID string
}

// TaskSelectedMsg indicates the active task selection.
type TaskSelectedMsg struct {
	TaskID string
}

// SubTaskSelectedMsg indicates the active subtask selection.
type SubTaskSelectedMsg struct {
	TaskID    string
	SubTaskID string
}

// ApiErrorMsg captures an API error scoped to a request.
type ApiErrorMsg struct {
	ReqID int64
	Err   error
	Scope string
}
