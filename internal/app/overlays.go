package app

import (
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/ui/components/toast"
)

// ToastMsg requests a toast notification to be shown.
type ToastMsg struct {
	Text  string
	Level toast.Level
}

// Overlays owns transient UI layers like toasts.
type Overlays struct {
	Toast toast.Model
}

// NewOverlays constructs the overlay state with defaults.
func NewOverlays() *Overlays {
	return &Overlays{
		Toast: toast.New(),
	}
}

// Update applies overlay-specific messages.
func (o *Overlays) Update(msg tea.Msg) tea.Cmd {
	if m, ok := msg.(ToastMsg); ok {
		o.Toast = o.Toast.Show(m.Text, m.Level)
	}

	return nil
}

// View renders the active overlays as a string.
func (o *Overlays) View() string {
	return o.Toast.View()
}
