package tasks

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/store"
)

// Model represents the tasks list screen state.
type Model struct {
	List list.Model
}

// New returns a tasks model with default list configuration.
func New() *Model {
	taskList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	taskList.Title = "Tasks"
	taskList.SetShowStatusBar(false)
	taskList.SetFilteringEnabled(false)
	taskList.SetShowHelp(false)
	return &Model{List: taskList}
}

// Init starts the initial task load for the tasks screen.
func (m *Model) Init(app screen.AppModel) tea.Cmd {
	reqID := app.NextReqID()
	return tea.Batch(
		func() tea.Msg { return store.TasksRequestedMsg{ReqID: reqID} },
		app.MeegleCmds().FetchTasks(app.ProjectKey(), reqID),
	)
}

func (m *Model) OnFocus(_ screen.AppModel) tea.Cmd {
	return nil
}

func (m *Model) OnBlur(_ screen.AppModel) {}

type taskItem struct {
	task store.Task
}

func (i taskItem) Title() string {
	return i.task.Name
}

func (i taskItem) Description() string {
	return i.task.ID
}

func (i taskItem) FilterValue() string {
	return i.task.Name
}

func (m *Model) syncItems(state store.State) {
	items := make([]list.Item, 0, len(state.Tasks))
	for _, task := range state.Tasks {
		items = append(items, taskItem{task: task})
	}
	m.List.SetItems(items)
}
