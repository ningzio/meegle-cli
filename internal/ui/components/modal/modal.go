package modal

import "github.com/charmbracelet/lipgloss"

type Model struct {
	Title   string
	Body    string
	Visible bool
}

func New() Model {
	return Model{}
}

func (m Model) View() string {
	if !m.Visible {
		return ""
	}

	box := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1, 2)
	return box.Render(m.Title + "\n\n" + m.Body)
}
