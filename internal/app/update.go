package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/store"
)

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if cmd, handled := a.handleGlobal(msg); handled {
		return a, cmd
	}

	a.Store = store.Reduce(a.Store, msg)

	screenCmd := a.Router.Current().Update(a, msg)
	overlayCmd := a.Overlays.Update(msg)

	return a, tea.Batch(screenCmd, overlayCmd)
}

func (a *App) handleGlobal(msg tea.Msg) (tea.Cmd, bool) {
	switch m := msg.(type) {
	case tea.KeyMsg:
		if a.KeyMap.Quit.Matches(m) {
			return tea.Quit, true
		}
		if a.KeyMap.Back.Matches(m) {
			return a.Router.Pop(a), true
		}
	}

	return nil, false
}
