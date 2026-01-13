package help

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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
	var content strings.Builder
	content.WriteString("Help")
	for _, line := range m.Lines {
		content.WriteString("\n" + line)
	}
	return style.Render(content.String())
}
