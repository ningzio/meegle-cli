package toast

import "github.com/charmbracelet/lipgloss"

// Level identifies the severity of a toast notification.
type Level string

const (
	// LevelInfo represents informational toasts.
	LevelInfo Level = "info"
	// LevelError represents error toasts.
	LevelError Level = "error"
	// LevelSuccess represents success toasts.
	LevelSuccess Level = "success"
)

// Model holds the state of a toast notification.
type Model struct {
	Message string
	Level   Level
	Visible bool
}

// New returns a hidden toast model.
func New() Model {
	return Model{}
}

// Show returns a visible toast populated with the given message.
func (m Model) Show(message string, level Level) Model {
	m.Message = message
	m.Level = level
	m.Visible = true
	return m
}

// View renders the toast if it is visible.
func (m Model) View() string {
	if !m.Visible || m.Message == "" {
		return ""
	}

	style := lipgloss.NewStyle().Padding(0, 1).Foreground(lipgloss.Color("231")).Background(lipgloss.Color("62"))
	if m.Level == LevelError {
		style = style.Background(lipgloss.Color("160"))
	}
	if m.Level == LevelSuccess {
		style = style.Background(lipgloss.Color("34"))
	}

	return style.Render(m.Message)
}
