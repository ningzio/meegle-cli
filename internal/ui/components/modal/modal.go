package modal

import "github.com/charmbracelet/lipgloss"

// Model represents a modal dialog.
type Model struct {
	Title   string
	Body    string
	Visible bool
}

// New returns a hidden modal model.
func New() Model {
	return Model{}
}

// View renders the modal when it is visible.
func (m Model) View() string {
	if !m.Visible {
		return ""
	}

	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
	return box.Render(m.Title + "\n\n" + m.Body)
}
