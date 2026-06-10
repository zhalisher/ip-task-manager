package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
)

type TaskFilter struct {
	Search   string
	Status   string
	Priority string
	Page     int
	Limit    int
}

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
	GetAll(ctx context.Context, userID uuid.UUID, filter TaskFilter) ([]*model.Task, error)
}
