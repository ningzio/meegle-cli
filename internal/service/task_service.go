package service

import (
	"context"

	"github.com/ningzio/meegle-cli/internal/adapter/api"
	"github.com/ningzio/meegle-cli/internal/model"
)

// TaskService handles business logic for tasks.
type TaskService struct {
	client api.Client
}

// NewTaskService creates a new TaskService.
func NewTaskService(client api.Client) *TaskService {
	return &TaskService{client: client}
}

// ListTasks retrieves tasks from the API.
func (s *TaskService) ListTasks(ctx context.Context) ([]model.Task, error) {
	return s.client.GetTasks(ctx)
}
