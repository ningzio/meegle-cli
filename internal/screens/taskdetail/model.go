package taskdetail

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"meegle-cli/internal/screen"
	"meegle-cli/internal/store"
)

type Model struct {
	List list.Model
}

func New() *Model {
	subtaskList := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	subtaskList.Title = "Subtasks"
	subtaskList.SetShowStatusBar(false)
	subtaskList.SetFilteringEnabled(false)
	subtaskList.SetShowHelp(false)
	return &Model{List: subtaskList}
}

func (m *Model) Init(app screen.AppModel) tea.Cmd {
	return nil
}

func (m *Model) OnFocus(app screen.AppModel) tea.Cmd {
	state := app.StoreState()
	if state.SelectedTaskID == "" {
		return nil
	}

	reqID := app.NextReqID()
	return tea.Batch(
		func() tea.Msg { return store.SubTasksRequestedMsg{ReqID: reqID, TaskID: state.SelectedTaskID} },
		app.MeegleCmds().FetchSubTasks(app.ProjectKey(), state.SelectedTaskID, reqID),
	)
}

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
	subTasks := state.SubTasksByTaskID[state.SelectedTaskID]
	items := make([]list.Item, 0, len(subTasks))
	for _, subTask := range subTasks {
		items = append(items, subTaskItem{subTask: subTask})
	}
	m.List.SetItems(items)
}
