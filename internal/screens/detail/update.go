package detail

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
	"meegle-cli/internal/screens/editor"
	"meegle-cli/internal/store"
)

func (s *Screen) Update(msg tea.Msg, appModel *app.Model) tea.Cmd {
	switch typed := msg.(type) {
	case tea.WindowSizeMsg:
		s.list.SetSize(typed.Width-4, typed.Height-10)
	case tea.KeyMsg:
		switch typed.String() {
		case app.KeyBack:
			return func() tea.Msg { return app.PopScreenMsg{} }
		case app.KeyAdd:
			return func() tea.Msg { return app.PushScreenMsg{Screen: editor.NewSubTaskScreen(s.taskID)} }
		case app.KeyToggle:
			item, ok := s.list.SelectedItem().(subTaskItem)
			if !ok {
				return func() tea.Msg { return app.ToastMsg{Text: "Select a subtask", Kind: app.ToastError} }
			}
			nextDone := !item.subTask.Done
			verb := "complete"
			if item.subTask.Done {
				verb = "reopen"
			}
			return func() tea.Msg {
				return app.ShowConfirmMsg{
					Title:       "Confirm",
					Body:        fmt.Sprintf("%s subtask '%s'?", verb, item.subTask.Name),
					ConfirmText: "y",
					CancelText:  "n",
					Danger:      item.subTask.Done,
					OnConfirm: toggleSubTaskMsg{
						taskID:    s.taskID,
						subTaskID: item.subTask.ID,
						done:      nextDone,
					},
				}
			}
		}
	case toggleSubTaskMsg:
		reqID := store.NextReqID(&appModel.Store, store.ReqToggleSubTask)
		return meegle.ToggleSubTaskDoneCmd(appModel.Client, reqID, typed.taskID, typed.subTaskID, typed.done)
	}

	if typed, ok := msg.(store.SubTaskCreatedMsg); ok {
		if typed.Err == nil && typed.TaskID == s.taskID {
			s.setItems(appModel)
		}
	}
	if typed, ok := msg.(store.SubTaskToggledMsg); ok {
		if typed.Err != nil && store.IsLatest(appModel.Store, store.ReqToggleSubTask, typed.ReqID) {
			return func() tea.Msg {
				return app.ToastMsg{Text: "Failed to update subtask: " + typed.Err.Error(), Kind: app.ToastError}
			}
		}
		if typed.Err == nil && typed.TaskID == s.taskID {
			s.setItems(appModel)
		}
	}

	var cmd tea.Cmd
	s.list, cmd = s.list.Update(msg)
	return cmd
}

func (s *Screen) setItems(appModel *app.Model) {
	items := make([]list.Item, 0)
	for _, sub := range store.SubTasksFor(appModel.Store, s.taskID) {
		items = append(items, subTaskItem{subTask: sub})
	}
	s.list.SetItems(items)
}
