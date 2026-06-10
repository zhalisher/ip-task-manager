package model

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	CategoryID  uuid.UUID
	Title       string
	Description string
	Status      string
	Priority    string
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
