package taskdetail

import (
	"fmt"
	"strings"

	"meegle-cli/internal/screen"
)

// View renders the task detail list or an empty-state message.
func (m *Model) View(app screen.AppModel) string {
	state := app.StoreState()
	if state.SelectedTaskID == "" {
		return "No task selected."
	}

	header := fmt.Sprintf("Task: %s", state.TasksByID[state.SelectedTaskID].Name)
	help := "a: add subtask  c: complete  r: rollback  esc: back"
	parts := []string{header, m.List.View(), help}
	return strings.Join(parts, "\n")
}
