package meegle

// Task describes a task returned by the Meegle API.
type Task struct {
	ID   string
	Name string
}

// SubTask describes a subtask returned by the Meegle API.
type SubTask struct {
	ID     string
	Name   string
	Status string
}
