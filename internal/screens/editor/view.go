package editor

import "meegle-cli/internal/screen"

func (m *Model) View(_ screen.AppModel) string {
	title := "Create Task"
	if m.Mode == ModeSubTask {
		title = "Create Subtask"
	}
	return title + "\n\n" + m.Input.View() + "\n\nenter: save  esc: cancel"
}
