package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
)

type UserUsecase interface {
	GetProfile(ctx context.Context, id uuid.UUID) (*model.User, error)
}
