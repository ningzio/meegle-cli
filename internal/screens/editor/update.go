package editor

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
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
				return func() tea.Msg { return app.TriggerCreateTaskMsg{Name: value} }
			}
			return func() tea.Msg { return app.TriggerCreateSubTaskMsg{TaskID: s.taskID, Name: value} }
		}
	}

	switch typed := msg.(type) {
	case store.TaskCreatedMsg:
		if s.mode == ModeTask {
			s.submitting = false
			if typed.Err == nil {
				return func() tea.Msg { return app.PopScreenMsg{} }
			}
		}
	case store.SubTaskCreatedMsg:
		if s.mode == ModeSubTask {
			s.submitting = false
			if typed.Err == nil {
				return func() tea.Msg { return app.PopScreenMsg{} }
			}
		}
	}

	var cmd tea.Cmd
	s.input, cmd = s.input.Update(msg)
	return cmd
}
