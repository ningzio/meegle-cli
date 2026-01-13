package api

import (
	"context"
	"testing"
)

func TestMockClient_GetTasks(t *testing.T) {
	client := NewMockClient()
	tasks, err := client.GetTasks(context.Background())
	if err != nil {
		t.Errorf("GetTasks() error = %v", err)
	}
	if len(tasks) == 0 {
		t.Errorf("GetTasks() returned empty list")
	}
}
