package detail

import (
	"strings"

	"meegle-cli/internal/app"
	"meegle-cli/internal/store"
)

func (s *Screen) View(appModel *app.Model) string {
	s.setItems(appModel)
	selected := store.SelectedTask(appModel.Store)
	title := "Task detail"
	if selected != nil {
		title = selected.Name
	}

	header := appModel.Theme.Title.Render("Task Detail")
	subtitle := appModel.Theme.Subtitle.Render(title)
	meta := appModel.Theme.Muted.Render("space:toggle • a:new subtask • esc:back")

	body := s.list.View()
	if len(store.SubTasksFor(appModel.Store, s.taskID)) == 0 {
		body = appModel.Theme.Muted.Render("No subtasks yet. Press 'a' to add one.")
	}
	footer := appModel.Theme.Footer.Render("q:quit")
	return strings.Join([]string{header, subtitle, meta, "", body, "", footer}, "\n")
}
