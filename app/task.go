package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/wbso/dtsgotask/store"
)

type Task struct {
	ID       uuid.UUID `json:"id"`
	Detail   string    `json:"detail"`
	IsDone   bool      `json:"done"`
	Assignee string    `json:"asignee"`
	Deadline time.Time `json:"deadline"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// List tasks
type ListTasksRequest struct {
	Show string `json:"show"`
}
type ListTasksResponse struct {
	Tasks []Task `json:"tasks"`
}

func (app *App) ListTasks(ctx context.Context, input ListTasksRequest) ListTasksResponse {
	tasks, err := app.Repo.ListTasks(ctx)
	if err != nil {
		return ListTasksResponse{}
	}
	return ListTasksResponse{Tasks: taskStoreListToTaskList(tasks)}
}

// Create a new task
type CreateTaskRequest struct {
	Detail   string `json:"detail"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
}

func (app *App) CreateTask(ctx context.Context, input CreateTaskRequest) (Task, error) {
	var output Task

	deadline, err := parseYMDToTime(input.Deadline)
	if err != nil {
		return output, err
	}
	taskStore, err := app.Repo.CreateTask(ctx, store.CreateTaskParams{
		ID:       uuid.New(),
		Detail:   input.Detail,
		Assignee: input.Assignee,
		Deadline: deadline,
	})
	if err != nil {
		return output, err
	}
	return taskStoreToTask(taskStore), nil
}

// Get a task by id
func (app *App) GetTask(ctx context.Context, id uuid.UUID) (Task, error) {
	taskStore, err := app.Repo.GetTaskById(ctx, id)
	if err != nil {
		return Task{}, err
	}
	return taskStoreToTask(taskStore), nil
}

// Update a task by id
type UpdateTaskRequest struct {
	Detail   string `json:"detail"`
	Assignee string `json:"assignee"`
	Deadline string `json:"deadline"`
}

func (app *App) UpdateTask(ctx context.Context, id uuid.UUID, input UpdateTaskRequest) (Task, error) {
	deadline, err := parseYMDToTime(input.Deadline)
	if err != nil {
		return Task{}, err
	}
	task, err := app.Repo.UpdateTask(ctx, store.UpdateTaskParams{
		ID:       id,
		Detail:   input.Detail,
		Assignee: input.Assignee,
		Deadline: deadline,
	})
	if err != nil {
		return Task{}, err
	}
	return Task(task), nil
}

// Delete a task by id
type DeleteTaskResponse struct {
	ID uuid.UUID `json:"id"`
}

func (app *App) DeleteTask(ctx context.Context, id uuid.UUID) DeleteTaskResponse {
	return DeleteTaskResponse{ID: id}
}

// Set a task as completed
type CompleteTaskResponse struct {
	ID uuid.UUID `json:"id"`
}

func (app *App) CompleteTask(ctx context.Context, id uuid.UUID) (Task, error) {
	var output Task
	taskStore, err := app.Repo.SetTaskDone(ctx, id)
	if err != nil {
		return output, err
	}
	return taskStoreToTask(taskStore), nil
}
