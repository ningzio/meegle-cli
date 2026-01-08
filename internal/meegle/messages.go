package meegle

// Store-facing messages emitted by Cmd factory functions.
type TasksFetchedMsg struct {
	ReqID int
	Tasks []Task
	Err   error
}

type TaskCreatedMsg struct {
	ReqID int
	Task  Task
	Err   error
}

type SubTaskCreatedMsg struct {
	ReqID   int
	TaskID  string
	SubTask SubTask
	Err     error
}

type SubTaskToggledMsg struct {
	ReqID   int
	TaskID  string
	SubTask SubTask
	Err     error
}
