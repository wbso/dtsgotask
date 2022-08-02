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
	Detail   string    `json:"detail"`
	Assignee string    `json:"assignee"`
	Deadline time.Time `json:"deadline"`
}

type CreateTaskResponse struct {
	ID uuid.UUID `json:"id"`
}

func (app *App) CreateTask(ctx context.Context, input CreateTaskRequest) (CreateTaskResponse, error) {
	var output CreateTaskResponse
	taskStore, err := app.Repo.CreateTask(ctx, store.CreateTaskParams{
		ID:       uuid.New(),
		Detail:   input.Detail,
		Assignee: input.Assignee,
		Deadline: input.Deadline,
	})
	if err != nil {
		return output, err
	}
	output.ID = taskStore.ID
	return output, nil
}

// Get a task by id
func (app *App) GetTask(ctx context.Context, id uuid.UUID) Task {
	return Task{}
}

// Update a task by id
type UpdateTaskRequest struct {
	Detail   string    `json:"description"`
	Asignee  string    `json:"asignee"`
	Deadline time.Time `json:"deadline"`
}
type UpdateTaskResponse struct {
	ID uuid.UUID `json:"id"`
}

func (app *App) UpdateTask(ctx context.Context, id uuid.UUID, input UpdateTaskRequest) UpdateTaskResponse {
	return UpdateTaskResponse{ID: id}
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

func (app *App) CompleteTask(ctx context.Context, id uuid.UUID) CompleteTaskResponse {
	return CompleteTaskResponse{ID: id}
}

func taskStoreListToTaskList(taskStores []store.Task) []Task {
	tasks := make([]Task, len(taskStores))
	for i, taskStore := range taskStores {
		tasks[i] = taskStoreToTask(taskStore)
	}
	return tasks
}

func taskStoreToTask(taskStore store.Task) Task {
	return Task{
		ID:        taskStore.ID,
		Detail:    taskStore.Detail,
		IsDone:    taskStore.IsDone,
		Assignee:  taskStore.Assignee,
		Deadline:  taskStore.Deadline,
		CreatedAt: taskStore.CreatedAt,
		UpdatedAt: taskStore.UpdatedAt,
	}
}
