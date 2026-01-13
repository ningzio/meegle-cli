package help

import "meegle-cli/internal/screen"

// View renders the help screen.
func (m *Model) View(_ screen.AppModel) string {
	return m.overlay.View()
}
