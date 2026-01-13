package taskdetail

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/screens/editor"
	"meegle-cli/internal/store"
)

func (m *Model) Update(app screen.AppModel, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case store.SubTasksLoadedMsg, store.SubTaskCreatedMsg, store.SubTaskCompletedMsg, store.SubTaskRolledBackMsg:
		m.syncItems(app.StoreState())
	case tea.WindowSizeMsg:
		m.List.SetSize(msg.Width, msg.Height-6)
	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			return app.Push(editor.NewSubTask())
		case "c":
			if item, ok := m.List.SelectedItem().(subTaskItem); ok {
				state := app.StoreState()
				return app.MeegleCmds().CompleteSubTask(app.ProjectKey(), state.SelectedTaskID, item.subTask.ID)
			}
		case "r":
			if item, ok := m.List.SelectedItem().(subTaskItem); ok {
				state := app.StoreState()
				return app.MeegleCmds().RollbackSubTask(app.ProjectKey(), state.SelectedTaskID, item.subTask.ID)
			}
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return cmd
}
