package form

import "github.com/charmbracelet/lipgloss"

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
	content := m.Title
	for _, field := range m.Fields {
		content += "\n" + field.Label + ": " + field.Value
	}
	return style.Render(content)
}
