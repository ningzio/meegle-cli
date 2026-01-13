package taskdetail_test

import (
	"testing"

	"github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"meegle-cli/internal/meegle"
	screenmock "meegle-cli/internal/screen/mock"
	"meegle-cli/internal/screens/taskdetail"
	"meegle-cli/internal/store"
)

func TestModelOnFocusUsesCachedSubTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	app := screenmock.NewMockAppModel(ctrl)

	state := store.State{
		SubTasksByTaskID: map[string][]store.SubTask{
			"task-1": {
				{ID: "sub-1", Name: "Draft plan", Status: "open"},
				{ID: "sub-2", Name: "Review details", Status: "completed"},
			},
		},
	}
	app.EXPECT().StoreState().Return(state)

	model := taskdetail.New("task-1")
	cmd := model.OnFocus(app)
	assert.Nil(t, cmd)

	items := model.List.Items()
	if assert.Len(t, items, 2) {
		first, ok := items[0].(interface{ Title() string })
		if assert.True(t, ok) {
			assert.Equal(t, "Draft plan", first.Title())
		}
		second, ok := items[1].(interface{ Title() string })
		if assert.True(t, ok) {
			assert.Equal(t, "Review details", second.Title())
		}
	}
}

func TestModelOnFocusFetchesSubTasksWhenEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	app := screenmock.NewMockAppModel(ctrl)

	state := store.State{
		SubTasksByTaskID: map[string][]store.SubTask{},
	}
	reqID := int64(7)
	projectKey := "proj-1"
	cmds := meegle.NewCmds(meegle.NewClient("https://example.com"), meegle.NewAuthManager("id", "secret", "user"))

	app.EXPECT().StoreState().Return(state)
	app.EXPECT().NextReqID().Return(reqID)
	app.EXPECT().MeegleCmds().Return(cmds)
	app.EXPECT().ProjectKey().Return(projectKey)

	model := taskdetail.New("task-1")
	cmd := model.OnFocus(app)
	assert.NotNil(t, cmd)

	msg := cmd()
	batch, ok := msg.(tea.BatchMsg)
	assert.True(t, ok)
	assert.Len(t, batch, 2)

	requestMsg, ok := batch[0]().(store.SubTasksRequestedMsg)
	assert.True(t, ok)
	assert.Equal(t, reqID, requestMsg.ReqID)
	assert.Equal(t, "task-1", requestMsg.TaskID)

	loadedMsg, ok := batch[1]().(store.SubTasksLoadedMsg)
	assert.True(t, ok)
	assert.Equal(t, reqID, loadedMsg.ReqID)
	assert.Equal(t, "task-1", loadedMsg.TaskID)
	assert.NotEmpty(t, loadedMsg.SubTasks)
}
