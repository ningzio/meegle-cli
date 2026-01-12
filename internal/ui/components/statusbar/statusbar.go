package statusbar

import "github.com/charmbracelet/lipgloss"

type Model struct {
	Left  string
	Right string
}

func New(left, right string) Model {
	return Model{Left: left, Right: right}
}

func (m Model) View(width int) string {
	style := lipgloss.NewStyle().Padding(0, 1)
	space := width - lipgloss.Width(m.Left) - lipgloss.Width(m.Right)
	if space < 1 {
		space = 1
	}

	return style.Render(m.Left + lipgloss.NewStyle().Width(space).Render(" ") + m.Right)
}
