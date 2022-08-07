package app

import (
	"context"

	"github.com/google/uuid"
)

// Delete a task by id

func (app *App) DeleteTask(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	err := app.Repo.DeleteTask(ctx, id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
