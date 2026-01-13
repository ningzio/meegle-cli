package app

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Header lipgloss.Style
	Footer lipgloss.Style
}

func DefaultTheme() Theme {
	return Theme{
		Header: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")),
		Footer: lipgloss.NewStyle().Foreground(lipgloss.Color("241")),
	}
}
