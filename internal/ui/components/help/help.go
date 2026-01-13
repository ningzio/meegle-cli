package help

import "github.com/charmbracelet/lipgloss"

// Model represents the help overlay state.
type Model struct {
	Visible bool
	Lines   []string
}

// New creates a help model with the provided lines.
func New(lines []string) Model {
	return Model{Lines: lines}
}

// View renders the help overlay when visible.
func (m Model) View() string {
	if !m.Visible {
		return ""
	}

	style := lipgloss.NewStyle().Padding(1, 2)
	content := "Help"
	for _, line := range m.Lines {
		content += "\n" + line
	}
	return style.Render(content)
}
