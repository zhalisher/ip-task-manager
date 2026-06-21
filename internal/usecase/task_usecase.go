package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
)

type TaskUsecase interface {
	Create(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
	GetAll(ctx context.Context, userID uuid.UUID, filter repository.TaskFilter) ([]*model.Task, error)
}
