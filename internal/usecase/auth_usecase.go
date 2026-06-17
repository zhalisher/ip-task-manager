package usecase

import (
	"context"

	"github.com/zhalisher/ip-task-manager/internal/domain/model"
)

type AuthUsecase interface {
	Register(ctx context.Context, email, password, name string) (*model.User, error)
	Login(ctx context.Context, email, password string) (accessToken, refreshToken string, err error)
}
