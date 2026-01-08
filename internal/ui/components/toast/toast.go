package toast

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Kind int

const (
	Info Kind = iota
	Error
)

type Model struct {
	Text    string
	Kind    Kind
	Visible bool
}

type hideMsg struct{}

func (m Model) Show(text string, kind Kind) (Model, tea.Cmd) {
	m.Text = text
	m.Kind = kind
	m.Visible = true
	return m, tea.Tick(3*time.Second, func(time.Time) tea.Msg { return hideMsg{} })
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case hideMsg:
		m.Visible = false
		m.Text = ""
	}
	return m, nil
}

func (m Model) View(style lipgloss.Style) string {
	if !m.Visible {
		return ""
	}
	return style.Render(m.Text)
}
