package tasks

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/app"
	"meegle-cli/internal/meegle"
	"meegle-cli/internal/screens/detail"
	"meegle-cli/internal/screens/editor"
	"meegle-cli/internal/store"
)

func (s *Screen) Update(msg tea.Msg, appModel *app.Model) tea.Cmd {
	switch typed := msg.(type) {
	case tea.WindowSizeMsg:
		s.list.SetSize(typed.Width-4, typed.Height-8)
	case tea.KeyMsg:
		switch typed.String() {
		case app.KeyNew:
			return func() tea.Msg { return app.PushScreenMsg{Screen: editor.NewTaskScreen()} }
		case app.KeyAdd:
			selected := appModel.Store.SelectedTaskID
			if selected == "" {
				return func() tea.Msg { return app.ToastMsg{Text: "Select a task first", Kind: app.ToastError} }
			}
			return func() tea.Msg { return app.PushScreenMsg{Screen: editor.NewSubTaskScreen(selected)} }
		case app.KeySubmit:
			if selected := appModel.Store.SelectedTaskID; selected != "" {
				return func() tea.Msg { return app.PushScreenMsg{Screen: detail.NewScreen(selected)} }
			}
		case app.KeyRefresh:
			return func() tea.Msg { return fetchTasksMsg{} }
		}
	case fetchTasksMsg:
		s.loading = true
		reqID := store.NextReqID(&appModel.Store, store.ReqFetchTasks)
		return meegle.FetchTasksCmd(appModel.Client, reqID)
	}

	if typed, ok := msg.(meegle.TasksFetchedMsg); ok {
		s.loading = false
		if typed.Err != nil && store.IsLatest(appModel.Store, store.ReqFetchTasks, typed.ReqID) {
			return func() tea.Msg {
				return app.ToastMsg{Text: "Failed to fetch tasks: " + typed.Err.Error(), Kind: app.ToastError}
			}
		}
		if typed.Err == nil {
			s.setItems(appModel)
		}
	}
	if typed, ok := msg.(meegle.TaskCreatedMsg); ok {
		if typed.Err == nil {
			s.setItems(appModel)
		}
	}

	var cmd tea.Cmd
	s.list, cmd = s.list.Update(msg)
	s.updateSelection(appModel)
	return cmd
}

func (s *Screen) setItems(appModel *app.Model) {
	items := make([]list.Item, 0, len(appModel.Store.Tasks))
	for _, task := range appModel.Store.Tasks {
		items = append(items, taskItem{task: task})
	}
	s.list.SetItems(items)
	s.updateSelection(appModel)
}

func (s *Screen) updateSelection(appModel *app.Model) {
	item, ok := s.list.SelectedItem().(taskItem)
	if ok {
		appModel.Store.SelectedTaskID = item.task.ID
	}
}

func (s *Screen) loadingView(appModel *app.Model) string {
	return fmt.Sprintf("%s %s", appModel.Theme.Tag.Render("loading"), appModel.Theme.Muted.Render("Fetching tasks..."))
}
