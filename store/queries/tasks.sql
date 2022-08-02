-- name: CreateTask :one
INSERT INTO tasks (id, detail, assignee, deadline)
VALUES (@id, @detail, @assignee, @deadline)
RETURNING *;
-- name: ListTasks :many
SELECT *
FROM tasks
order by created_at desc;
-- name: GetTaskById :one
SELECT *
FROM tasks
WHERE id = @id;
-- name: UpdateTask :one
UPDATE tasks
SET detail = @detail,
    assignee = @assignee,
    deadline = @deadline,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;
-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = @id;
-- name: SetTaskDone :one
UPDATE tasks
SET is_done = true,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;
-- name: SetTaskUndone :one
UPDATE tasks
SET is_done = false,
    updated_at = CURRENT_TIMESTAMP
WHERE id = @id
RETURNING *;