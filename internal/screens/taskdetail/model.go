package taskdetail

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/store"
)

// Model represents the task detail screen state.
type Model struct {
	List           list.Model
	SelectedTaskID string
}

// New returns a task detail model with default list configuration.
func New(taskID string) *Model {
	subtaskList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	subtaskList.Title = "Subtasks"
	subtaskList.SetShowStatusBar(false)
	subtaskList.SetFilteringEnabled(false)
	subtaskList.SetShowHelp(false)
	return &Model{List: subtaskList, SelectedTaskID: taskID}
}

// Init prepares the task detail screen for first render.
func (m *Model) Init(app screen.AppModel) tea.Cmd {
	return nil
}

// OnFocus refreshes subtasks when the screen gains focus.
func (m *Model) OnFocus(app screen.AppModel) tea.Cmd {
	taskID := m.taskID(app.StoreState())
	if taskID == "" {
		return nil
	}

	reqID := app.NextReqID()
	return tea.Batch(
		func() tea.Msg { return store.SubTasksRequestedMsg{ReqID: reqID, TaskID: taskID} },
		app.MeegleCmds().FetchSubTasks(app.ProjectKey(), taskID, reqID),
	)
}

// OnBlur handles cleanup when the screen loses focus.
func (m *Model) OnBlur(app screen.AppModel) {}

type subTaskItem struct {
	subTask store.SubTask
}

func (i subTaskItem) Title() string {
	return i.subTask.Name
}

func (i subTaskItem) Description() string {
	return i.subTask.Status
}

func (i subTaskItem) FilterValue() string {
	return i.subTask.Name
}

func (m *Model) syncItems(state store.State) {
	taskID := m.taskID(state)
	subTasks := state.SubTasksByTaskID[taskID]
	items := make([]list.Item, 0, len(subTasks))
	for _, subTask := range subTasks {
		items = append(items, subTaskItem{subTask: subTask})
	}
	m.List.SetItems(items)
}

func (m *Model) taskID(state store.State) string {
	if m.SelectedTaskID != "" {
		return m.SelectedTaskID
	}
	return state.SelectedTaskID
}
