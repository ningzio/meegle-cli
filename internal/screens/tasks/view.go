package tasks

import (
	"meegle-cli/internal/screen"
)

func (m *Model) View(app screen.AppModel) string {
	state := app.StoreState()
	if len(state.Tasks) == 0 {
		return "No tasks loaded yet."
	}

	return m.List.View()
}
