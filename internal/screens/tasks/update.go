package tasks

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/screens/editor"
	"meegle-cli/internal/screens/taskdetail"
	"meegle-cli/internal/store"
)

// Update handles list interactions and navigation for the tasks screen.
func (m *Model) Update(app screen.AppModel, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case store.TasksLoadedMsg, store.TaskCreatedMsg:
		m.syncItems(app.StoreState())
	case tea.WindowSizeMsg:
		m.List.SetSize(msg.Width, msg.Height-4)
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item, ok := m.List.SelectedItem().(taskItem); ok {
				return tea.Sequence(
					func() tea.Msg { return store.TaskSelectedMsg{TaskID: item.task.ID} },
					app.Push(taskdetail.New()),
				)
			}
		case "n":
			return app.Push(editor.NewTask())
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return cmd
}
