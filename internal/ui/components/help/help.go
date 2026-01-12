package help

import "github.com/charmbracelet/lipgloss"

type Model struct {
	Visible bool
	Lines   []string
}

func New(lines []string) Model {
	return Model{Lines: lines}
}

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
