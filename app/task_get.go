package app

import (
	"context"

	"github.com/google/uuid"
)

// Get a task by id
func (app *App) GetTask(ctx context.Context, id uuid.UUID) (Task, error) {
	taskStore, err := app.Repo.GetTaskById(ctx, id)
	if err != nil {
		return Task{}, err
	}
	return taskStoreToTask(taskStore), nil
}
