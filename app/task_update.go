package app

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wbso/dtsgotask/store"
)

type UpdateTaskRequest struct {
	Detail   string `json:"detail"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
}

func (r *UpdateTaskRequest) Validate() error {
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

func (app *App) UpdateTask(ctx context.Context, id uuid.UUID, input UpdateTaskRequest) (Task, error) {
	if err := input.Validate(); err != nil {
		return Task{}, err
	}
	deadline, err := parseYMDToTime(input.Deadline)
	if err != nil {
		return Task{}, InvalidInputError("invalid date format")
	}
	task, err := app.Repo.UpdateTask(ctx, store.UpdateTaskParams{
		ID:       id,
		Detail:   input.Detail,
		Assignee: input.Assignee,
		Deadline: deadline,
	})
	if err != nil {
		return Task{}, fmt.Errorf("failed to update task: %w", err)
	}
	return Task(task), nil
}

// Set a task as completed
type CompleteTaskResponse struct {
	ID uuid.UUID `json:"id"`
}

func (app *App) SetTaskDone(ctx context.Context, id uuid.UUID) (Task, error) {
	var output Task
	taskStore, err := app.Repo.SetTaskDone(ctx, id)
	if err != nil {
		return output, err
	}
	return taskStoreToTask(taskStore), nil
}

func (app *App) SetTaskUndone(ctx context.Context, id uuid.UUID) (Task, error) {
	var output Task
	taskStore, err := app.Repo.SetTaskUndone(ctx, id)
	if err != nil {
		return output, err
	}
	return taskStoreToTask(taskStore), nil
}
