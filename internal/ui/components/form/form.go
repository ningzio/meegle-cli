package form

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Title   string
	Fields  []Field
	Focused int
}

type Field struct {
	Label string
	Value string
}

func New(title string) Model {
	return Model{Title: title}
}

func (m Model) View() string {
	style := lipgloss.NewStyle().Padding(1, 2)
	var content strings.Builder
	content.WriteString(m.Title)
	for _, field := range m.Fields {
		content.WriteString("\n" + field.Label + ": " + field.Value)
	}
	return style.Render(content.String())
}
