package modal

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Visible     bool
	Title       string
	Body        string
	ConfirmText string
	CancelText  string
	Danger      bool
	OnConfirm   tea.Msg
}

func (m Model) Show(title, body, confirmText, cancelText string, danger bool, onConfirm tea.Msg) Model {
	m.Visible = true
	m.Title = title
	m.Body = body
	m.ConfirmText = confirmText
	m.CancelText = cancelText
	m.Danger = danger
	m.OnConfirm = onConfirm
	return m
}

func (m Model) Hide() Model {
	m.Visible = false
	return m
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.Visible {
		return m, nil
	}
	switch typed := msg.(type) {
	case tea.KeyMsg:
		switch typed.String() {
		case "y", "enter":
			m.Visible = false
			return m, func() tea.Msg { return m.OnConfirm }
		case "n", "esc":
			m.Visible = false
			return m, nil
		}
	}
	return m, nil
}

func (m Model) View(theme Theme) string {
	if !m.Visible {
		return ""
	}
	lines := []string{
		theme.Title.Render(m.Title),
		theme.Body.Render(m.Body),
		"",
		strings.Join([]string{
			theme.Button.Render(m.ConfirmText),
			theme.Muted.Render(" / "),
			theme.Muted.Render(m.CancelText),
		}, ""),
	}
	return theme.Frame.Render(strings.Join(lines, "\n"))
}

type Theme struct {
	Frame  lipgloss.Style
	Title  lipgloss.Style
	Body   lipgloss.Style
	Button lipgloss.Style
	Muted  lipgloss.Style
}
