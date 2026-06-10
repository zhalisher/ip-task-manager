package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
