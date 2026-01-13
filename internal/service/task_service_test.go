package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ningzio/meegle-cli/internal/model"
)

type MockClient struct {
	Tasks []model.Task
	Err   error
}

func (m *MockClient) GetTasks(ctx context.Context) ([]model.Task, error) {
	return m.Tasks, m.Err
}

func TestListTasks(t *testing.T) {
	tests := []struct {
		name    string
		tasks   []model.Task
		err     error
		wantErr bool
	}{
		{
			name: "success",
			tasks: []model.Task{
				{ID: "1", Title: "Test Task"},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:    "error",
			tasks:   nil,
			err:     errors.New("api error"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &MockClient{Tasks: tt.tasks, Err: tt.err}
			svc := NewTaskService(client)
			got, err := svc.ListTasks(context.Background())

			if (err != nil) != tt.wantErr {
				t.Errorf("ListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && len(got) != len(tt.tasks) {
				t.Errorf("ListTasks() got = %v, want %v", got, tt.tasks)
			}
		})
	}
}
