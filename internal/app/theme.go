package app

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	Title       lipgloss.Style
	Subtitle    lipgloss.Style
	Selected    lipgloss.Style
	Muted       lipgloss.Style
	Tag         lipgloss.Style
	Footer      lipgloss.Style
	Border      lipgloss.Style
	Modal       lipgloss.Style
	ModalTitle  lipgloss.Style
	ModalButton lipgloss.Style
	ModalDanger lipgloss.Style
	Toast       lipgloss.Style
	ToastError  lipgloss.Style
}

func DefaultTheme() Theme {
	accent := lipgloss.Color("86")
	muted := lipgloss.Color("243")
	errorColor := lipgloss.Color("203")
	return Theme{
		Title:       lipgloss.NewStyle().Foreground(accent).Bold(true),
		Subtitle:    lipgloss.NewStyle().Foreground(lipgloss.Color("69")).Bold(true),
		Selected:    lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(accent).Bold(true),
		Muted:       lipgloss.NewStyle().Foreground(muted),
		Tag:         lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(lipgloss.Color("62")).Padding(0, 1),
		Footer:      lipgloss.NewStyle().Foreground(muted),
		Border:      lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(muted).Padding(0, 1),
		Modal:       lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(accent).Padding(1, 2).Background(lipgloss.Color("235")),
		ModalTitle:  lipgloss.NewStyle().Foreground(accent).Bold(true),
		ModalButton: lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(accent).Padding(0, 1),
		ModalDanger: lipgloss.NewStyle().Foreground(lipgloss.Color("230")).Background(errorColor).Padding(0, 1),
		Toast:       lipgloss.NewStyle().Foreground(lipgloss.Color("229")).Background(lipgloss.Color("35")).Padding(0, 1),
		ToastError:  lipgloss.NewStyle().Foreground(lipgloss.Color("229")).Background(errorColor).Padding(0, 1),
	}
}
