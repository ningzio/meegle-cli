package app

import "github.com/charmbracelet/lipgloss"

// Theme defines shared styles for the application chrome.
type Theme struct {
	Header lipgloss.Style
	Footer lipgloss.Style
}

// DefaultTheme returns the default styling used by the app.
func DefaultTheme() Theme {
	return Theme{
		Header: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63")),
		Footer: lipgloss.NewStyle().Foreground(lipgloss.Color("241")),
	}
}
