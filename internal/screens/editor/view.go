package editor

import (
	"strings"

	"meegle-cli/internal/app"
)

func (s *Screen) View(appModel *app.Model) string {
	title := "New Task"
	help := "enter:save • esc:back"
	if s.mode == ModeSubTask {
		title = "New Subtask"
	}
	header := appModel.Theme.Title.Render(title)
	body := s.input.View()
	if s.submitting {
		body = appModel.Theme.Muted.Render("Saving...")
	}
	footer := appModel.Theme.Footer.Render(help)
	return strings.Join([]string{header, "", body, "", footer}, "\n")
}
