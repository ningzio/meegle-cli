package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/ui/components/toast"
)

type ToastMsg struct {
	Text  string
	Level toast.Level
}

type Overlays struct {
	Toast toast.Model
}

func NewOverlays() *Overlays {
	return &Overlays{
		Toast: toast.New(),
	}
}

func (o *Overlays) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case ToastMsg:
		o.Toast = o.Toast.Show(m.Text, m.Level)
	}

	return nil
}

func (o *Overlays) View() string {
	return o.Toast.View()
}
