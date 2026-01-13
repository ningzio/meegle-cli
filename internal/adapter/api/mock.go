package api

import (
	"context"

	"github.com/ningzio/meegle-cli/internal/model"
)

// MockClient is a mock implementation of the Client interface.
type MockClient struct{}

// NewMockClient creates a new MockClient.
func NewMockClient() *MockClient {
	return &MockClient{}
}

// GetTasks returns a fixed list of tasks for testing/development.
func (m *MockClient) GetTasks(ctx context.Context) ([]model.Task, error) {
	return []model.Task{
		{ID: "1", Title: "Setup Project", Description: "Initialize the repository", Status: "Done"},
		{ID: "2", Title: "Implement TUI", Description: "Create Bubble Tea models", Status: "In Progress"},
	}, nil
}
