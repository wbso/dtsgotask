-- name: CreateTask :one
INSERT INTO tasks (id, detail, assignee, deadline)
VALUES (@id, @detail, @assignee, @deadline)
RETURNING *;
-- name: ListTasks :many
SELECT *
FROM tasks
order by created_at desc;