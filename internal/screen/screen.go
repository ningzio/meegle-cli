package screen

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/meegle"
	"meegle-cli/internal/store"
)

type AppModel interface {
	StoreState() store.State
	MeegleCmds() *meegle.Cmds
	ProjectKey() string
	NextReqID() int64
	Push(screen Screen) tea.Cmd
	Pop() tea.Cmd
	Replace(screen Screen) tea.Cmd
}

type Screen interface {
	Init(app AppModel) tea.Cmd
	Update(app AppModel, msg tea.Msg) tea.Cmd
	View(app AppModel) string

	OnFocus(app AppModel) tea.Cmd
	OnBlur(app AppModel)
}
