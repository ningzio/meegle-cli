package form

import "github.com/charmbracelet/lipgloss"

// Model represents a labeled form with multiple fields.
type Model struct {
	Title   string
	Fields  []Field
	Focused int
}

// Field represents a labeled value inside a form.
type Field struct {
	Label string
	Value string
}

// New creates a form model with the provided title.
func New(title string) Model {
	return Model{Title: title}
}

// View renders the form as a string.
func (m Model) View() string {
	style := lipgloss.NewStyle().Padding(1, 2)
	content := m.Title
	for _, field := range m.Fields {
		content += "\n" + field.Label + ": " + field.Value
	}
	return style.Render(content)
}
