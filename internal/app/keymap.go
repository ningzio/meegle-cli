package app

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines the key bindings used by the application.
type KeyMap struct {
	Quit key.Binding
	Back key.Binding
}

// DefaultKeyMap returns the default key bindings for navigation.
func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "back"),
		),
	}
}
