package app

import tea "github.com/charmbracelet/bubbletea"

type PushScreenMsg struct{ Screen Screen }

type PopScreenMsg struct{}

type GoToScreenMsg struct{ Screen Screen }

type ToastMsg struct {
	Text string
	Kind ToastKind
}

type ToastKind int

const (
	ToastInfo ToastKind = iota
	ToastError
)

type ShowConfirmMsg struct {
	Title       string
	Body        string
	ConfirmText string
	CancelText  string
	Danger      bool
	OnConfirm   tea.Msg
}

type TriggerFetchTasksMsg struct{}

type TriggerCreateTaskMsg struct {
	Name string
}

type TriggerCreateSubTaskMsg struct {
	TaskID string
	Name   string
}

type TriggerToggleSubTaskMsg struct {
	TaskID    string
	SubTaskID string
	Done      bool
}
