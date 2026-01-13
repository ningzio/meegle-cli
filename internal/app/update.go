package app

import (
	"github.com/charmbracelet/bubbles/key"
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
	if m, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(m, a.KeyMap.Quit) {
			return tea.Quit, true
		}
		if key.Matches(m, a.KeyMap.Back) {
			return a.Pop(), true
		}
	}

	return nil, false
}
