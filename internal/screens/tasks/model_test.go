package tasks_test

import (
	"reflect"
	"testing"

	"github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"meegle-cli/internal/meegle"
	screenmock "meegle-cli/internal/screen/mock"
	"meegle-cli/internal/screens/editor"
	"meegle-cli/internal/screens/taskdetail"
	"meegle-cli/internal/screens/tasks"
	"meegle-cli/internal/store"
)

type pushMsg struct {
	value string
}

func TestModelInit(t *testing.T) {
	testCases := []struct {
		name   string
		reqID  int64
		projID string
	}{
		{name: "builds batch commands", reqID: 42, projID: "proj-1"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			app := screenmock.NewMockAppModel(ctrl)

			cmds := meegle.NewCmds(meegle.NewClient("https://example.com"), meegle.NewAuthManager("id", "secret", "user"))
			app.EXPECT().NextReqID().Return(tc.reqID)
			app.EXPECT().MeegleCmds().Return(cmds)
			app.EXPECT().ProjectKey().Return(tc.projID)

			model := tasks.New()
			cmd := model.Init(app)

			assert.NotNil(t, cmd)

			msg := cmd()
			batch, ok := msg.(tea.BatchMsg)
			assert.True(t, ok)
			assert.Len(t, batch, 2)

			requestMsg, ok := batch[0]().(store.TasksRequestedMsg)
			assert.True(t, ok)
			assert.Equal(t, tc.reqID, requestMsg.ReqID)

			loadedMsg, ok := batch[1]().(store.TasksLoadedMsg)
			assert.True(t, ok)
			assert.Equal(t, tc.reqID, loadedMsg.ReqID)
			assert.NotEmpty(t, loadedMsg.Tasks)
		})
	}
}

func TestModelUpdateSyncItems(t *testing.T) {
	testCases := []struct {
		name       string
		msg        tea.Msg
		state      store.State
		wantTitles []string
	}{
		{
			name: "tasks loaded syncs items",
			msg:  store.TasksLoadedMsg{ReqID: 1},
			state: store.State{
				Tasks: []store.Task{
					{ID: "task-1", Name: "Design Milestone"},
					{ID: "task-2", Name: "Build TUI"},
				},
			},
			wantTitles: []string{"Design Milestone", "Build TUI"},
		},
		{
			name: "task created syncs items",
			msg:  store.TaskCreatedMsg{Task: store.Task{ID: "task-3", Name: "Ship MVP"}},
			state: store.State{
				Tasks: []store.Task{
					{ID: "task-1", Name: "Design Milestone"},
					{ID: "task-2", Name: "Build TUI"},
					{ID: "task-3", Name: "Ship MVP"},
				},
			},
			wantTitles: []string{"Design Milestone", "Build TUI", "Ship MVP"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			app := screenmock.NewMockAppModel(ctrl)
			app.EXPECT().StoreState().Return(tc.state)

			model := tasks.New()
			_ = model.Update(app, tc.msg)

			items := model.List.Items()
			assert.Len(t, items, len(tc.wantTitles))
			for i, title := range tc.wantTitles {
				item, ok := items[i].(interface{ Title() string })
				if assert.True(t, ok) {
					assert.Equal(t, title, item.Title())
				}
			}
		})
	}
}

func TestModelUpdateKeyActions(t *testing.T) {
	testCases := []struct {
		name           string
		key            tea.KeyMsg
		setupModel     func(*tasks.Model, *screenmock.MockAppModel)
		expectBatch    bool
		expectSelected bool
	}{
		{
			name: "enter selects task and pushes detail screen",
			key:  tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(model *tasks.Model, app *screenmock.MockAppModel) {
				state := store.State{Tasks: []store.Task{{ID: "task-1", Name: "Design Milestone"}}}
				app.EXPECT().StoreState().Return(state)
				_ = model.Update(app, store.TasksLoadedMsg{ReqID: 1})
			},
			expectBatch:    true,
			expectSelected: true,
		},
		{
			name:        "n opens new task editor",
			key:         tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'n'}},
			setupModel:  func(_ *tasks.Model, _ *screenmock.MockAppModel) {},
			expectBatch: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			app := screenmock.NewMockAppModel(ctrl)

			model := tasks.New()
			tc.setupModel(model, app)

			pushCmd := func() tea.Msg { return pushMsg{value: "pushed"} }
			if tc.expectSelected {
				app.EXPECT().Push(gomock.AssignableToTypeOf(&taskdetail.Model{})).Return(pushCmd)
			} else {
				app.EXPECT().Push(gomock.AssignableToTypeOf(&editor.Model{})).Return(pushCmd)
			}

			cmd := model.Update(app, tc.key)
			assert.NotNil(t, cmd)

			msg := cmd()
			if tc.expectBatch {
				assert.Equal(t, "tea.sequenceMsg", reflect.TypeOf(msg).String())
				seqValue := reflect.ValueOf(msg)
				assert.Equal(t, 2, seqValue.Len())

				firstCmd, ok := seqValue.Index(0).Interface().(tea.Cmd)
				assert.True(t, ok)
				secondCmd, ok := seqValue.Index(1).Interface().(tea.Cmd)
				assert.True(t, ok)

				selectedMsg, ok := firstCmd().(store.TaskSelectedMsg)
				assert.True(t, ok)
				assert.Equal(t, "task-1", selectedMsg.TaskID)
				assert.IsType(t, pushMsg{}, secondCmd())
			} else {
				assert.IsType(t, pushMsg{}, msg)
			}
		})
	}
}
