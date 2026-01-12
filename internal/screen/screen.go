package screen

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/store"
)

type AppModel interface {
	StoreState() store.State
}

type Screen interface {
	Init(app AppModel) tea.Cmd
	Update(app AppModel, msg tea.Msg) tea.Cmd
	View(app AppModel) string

	OnFocus(app AppModel) tea.Cmd
	OnBlur(app AppModel)
}
