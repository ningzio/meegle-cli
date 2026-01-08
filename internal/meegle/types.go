package meegle

type Task struct {
	ID   string
	Name string
}

type SubTask struct {
	ID     string
	TaskID string
	Name   string
	Done   bool
}
