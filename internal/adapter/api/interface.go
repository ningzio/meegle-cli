package api

import (
	"context"

	"github.com/ningzio/meegle-cli/internal/model"
)

// Client defines the interface for the Meegle API.
type Client interface {
	GetTasks(ctx context.Context) ([]model.Task, error)
}
