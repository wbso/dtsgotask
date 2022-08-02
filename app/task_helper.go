package app

import "github.com/wbso/dtsgotask/store"

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
