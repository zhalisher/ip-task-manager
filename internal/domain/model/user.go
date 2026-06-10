package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
