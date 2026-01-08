package tasks

import (
	"fmt"
	"strings"

	"meegle-cli/internal/app"
)

func (s *Screen) View(appModel *app.Model) string {
	header := appModel.Theme.Title.Render("Meegle Tasks")
	sub := appModel.Theme.Muted.Render("n:new task • a:new subtask • enter:details • r:refresh")

	body := s.list.View()
	if s.loading {
		body = s.loadingView(appModel)
	}

	footer := appModel.Theme.Footer.Render("q:quit")
	return strings.Join([]string{header, sub, "", body, "", footer}, "\n")
}

func (s *Screen) Debug(appModel *app.Model) string {
	return fmt.Sprintf("selected=%s", appModel.Store.SelectedTaskID)
}
