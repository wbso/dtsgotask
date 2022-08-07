package app

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID       uuid.UUID `json:"id"`
	Detail   string    `json:"detail"`
	IsDone   bool      `json:"done"`
	Assignee string    `json:"assignee"`
	Deadline time.Time `json:"deadline"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
