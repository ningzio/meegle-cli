package toast

import "github.com/charmbracelet/lipgloss"

type Level string

const (
	LevelInfo    Level = "info"
	LevelError   Level = "error"
	LevelSuccess Level = "success"
)

type Model struct {
	Message string
	Level   Level
	Visible bool
}

func New() Model {
	return Model{}
}

func (m Model) Show(message string, level Level) Model {
	m.Message = message
	m.Level = level
	m.Visible = true
	return m
}

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
