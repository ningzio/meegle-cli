package help

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

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
	var content strings.Builder
	content.WriteString("Help")
	for _, line := range m.Lines {
		content.WriteString("\n")
		content.WriteString(line)
	}
	return style.Render(content.String())
}
