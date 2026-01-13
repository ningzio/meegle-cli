package editor

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/store"
)

func (m *Model) Update(app screen.AppModel, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			name := strings.TrimSpace(m.Input.Value())
			if name == "" {
				return nil
			}
			if m.Mode == ModeTask {
				return tea.Batch(
					app.MeegleCmds().CreateTask(app.ProjectKey(), name),
					app.Pop(),
				)
			}

			state := app.StoreState()
			if state.SelectedTaskID == "" {
				return app.Pop()
			}
			return tea.Batch(
				app.MeegleCmds().CreateSubTask(app.ProjectKey(), state.SelectedTaskID, name),
				app.Pop(),
			)
		case "esc":
			return app.Pop()
		}
	case store.TaskCreatedMsg, store.SubTaskCreatedMsg:
		return nil
	}

	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return cmd
}
