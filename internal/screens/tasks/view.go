package tasks

import (
	"strings"

	"meegle-cli/internal/app"
)

func (m *Model) View(app *app.App) string {
	lines := []string{"Tasks"}
	if len(app.Store.Tasks) == 0 {
		lines = append(lines, "No tasks loaded yet.")
	}
	for _, task := range app.Store.Tasks {
		lines = append(lines, "• "+task.Name)
	}
	return strings.Join(lines, "\n")
}
