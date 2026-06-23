package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) GetProfile(ctx context.Context, id uuid.UUID) (*model.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
