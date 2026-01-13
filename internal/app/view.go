package app

import "strings"

// View renders the application chrome and current screen output.
func (a *App) View() string {
	parts := []string{
		a.Theme.Header.Render("Meegle TUI"),
		a.Router.Current().View(a),
		a.Overlays.View(),
		a.Theme.Footer.Render("Press ? for help"),
	}

	return strings.Join(parts, "\n")
}
