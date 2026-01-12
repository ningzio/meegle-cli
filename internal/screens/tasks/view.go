package tasks

import (
	"strings"

	"meegle-cli/internal/screen"
)

func (m *Model) View(app screen.AppModel) string {
	lines := []string{"Tasks"}
	state := app.StoreState()
	if len(state.Tasks) == 0 {
		lines = append(lines, "No tasks loaded yet.")
	}
	for _, task := range state.Tasks {
		lines = append(lines, "• "+task.Name)
	}
	return strings.Join(lines, "\n")
}
