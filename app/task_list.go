package app

import "context"

func (app *App) ListTasks(ctx context.Context) ([]Task, error) {
	tasks, err := app.Repo.ListTasks(ctx)
	if err != nil {
		return []Task{}, err
	}
	return taskStoreListToTaskList(tasks), nil
}
