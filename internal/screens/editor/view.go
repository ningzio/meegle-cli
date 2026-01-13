package editor

import "meegle-cli/internal/screen"

// View renders the editor prompt and input control.
func (m *Model) View(app screen.AppModel) string {
	title := "Create Task"
	if m.Mode == ModeSubTask {
		title = "Create Subtask"
	}
	return title + "\n\n" + m.Input.View() + "\n\nenter: save  esc: cancel"
}
