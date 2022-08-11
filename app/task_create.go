package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/wbso/dtsgotask/store"
)

// Create a new task
type CreateTaskRequest struct {
	Detail   string `json:"detail"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Detail == "" {
		return InvalidInputError("detail is required")
	}
	if r.Assignee == "" {
		return InvalidInputError("assignee is required")
	}
	if r.Deadline == "" {
		return InvalidInputError("deadline is required")
	}
	return nil
}

func (app *App) CreateTask(ctx context.Context, input CreateTaskRequest) (Task, error) {
	if err := input.Validate(); err != nil {
		return Task{}, err
	}

	deadline, err := parseYMDToTime(input.Deadline)
	if err != nil {
		return Task{}, err
	}
	taskStore, err := app.Repo.CreateTask(ctx, store.CreateTaskParams{
		ID:       uuid.New(),
		Detail:   input.Detail,
		Assignee: input.Assignee,
		Deadline: deadline,
	})
	if err != nil {
		return Task{}, err
	}
	return taskStoreToTask(taskStore), nil
}
