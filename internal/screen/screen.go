package screen

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/meegle"
	"meegle-cli/internal/store"
)

// AppModel exposes the app services that screens can consume.
type AppModel interface {
	StoreState() store.State
	MeegleCmds() *meegle.Cmds
	ProjectKey() string
	NextReqID() int64
	Push(screen Screen) tea.Cmd
	Pop() tea.Cmd
	Replace(screen Screen) tea.Cmd
}

// Screen represents a navigable view in the TUI.
type Screen interface {
	Init(app AppModel) tea.Cmd
	Update(app AppModel, msg tea.Msg) tea.Cmd
	View(app AppModel) string

	OnFocus(app AppModel) tea.Cmd
	OnBlur(app AppModel)
}
