package editor

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
	"meegle-cli/internal/store"
)

func (s *Screen) Update(msg tea.Msg, appModel *app.Model) tea.Cmd {
	switch typed := msg.(type) {
	case tea.KeyMsg:
		switch typed.String() {
		case app.KeyBack:
			return func() tea.Msg { return app.PopScreenMsg{} }
		case app.KeySubmit:
			value := strings.TrimSpace(s.input.Value())
			if value == "" {
				return func() tea.Msg { return app.ToastMsg{Text: "Name cannot be empty", Kind: app.ToastError} }
			}
			s.submitting = true
			if s.mode == ModeTask {
				reqID := store.NextReqID(&appModel.Store, store.ReqCreateTask)
				return meegle.CreateTaskCmd(appModel.Client, reqID, value)
			}
			reqID := store.NextReqID(&appModel.Store, store.ReqCreateSubTask)
			return meegle.CreateSubTaskCmd(appModel.Client, reqID, s.taskID, value)
		}
	}

	switch typed := msg.(type) {
	case meegle.TaskCreatedMsg:
		if s.mode == ModeTask {
			s.submitting = false
			if typed.Err != nil && store.IsLatest(appModel.Store, store.ReqCreateTask, typed.ReqID) {
				return func() tea.Msg {
					return app.ToastMsg{Text: "Failed to create task: " + typed.Err.Error(), Kind: app.ToastError}
				}
			}
			if typed.Err == nil {
				return func() tea.Msg { return app.PopScreenMsg{} }
			}
		}
	case meegle.SubTaskCreatedMsg:
		if s.mode == ModeSubTask {
			s.submitting = false
			if typed.Err != nil && store.IsLatest(appModel.Store, store.ReqCreateSubTask, typed.ReqID) {
				return func() tea.Msg {
					return app.ToastMsg{Text: "Failed to create subtask: " + typed.Err.Error(), Kind: app.ToastError}
				}
			}
			if typed.Err == nil {
				return func() tea.Msg { return app.PopScreenMsg{} }
			}
		}
	}

	var cmd tea.Cmd
	s.input, cmd = s.input.Update(msg)
	return cmd
}
