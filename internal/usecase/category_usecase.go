package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
)

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository) *categoryUsecase {
	return &categoryUsecase{categoryRepo: categoryRepo}
}

func (u *categoryUsecase) Create(ctx context.Context, category *model.Category) error {
	return u.categoryRepo.Create(ctx, category)
}

func (u *categoryUsecase) Update(ctx context.Context, category *model.Category) error {
	_, err := u.categoryRepo.GetByID(ctx, category.ID)
	if err != nil {
		return errors.New("category not found")
	}
	return u.categoryRepo.Update(ctx, category)
}

func (u *categoryUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := u.categoryRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("category not found")
	}
	return u.categoryRepo.Delete(ctx, id)
}

func (u *categoryUsecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Category, error) {
	return u.categoryRepo.GetByID(ctx, id)
}

func (u *categoryUsecase) GetAll(ctx context.Context, userID uuid.UUID) ([]*model.Category, error) {
	return u.categoryRepo.GetAll(ctx, userID)
}
