package usecase

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/domain/model"
	"github.com/zhalisher/ip-task-manager/internal/domain/repository"
)

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) *taskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (u *taskUsecase) Create(ctx context.Context, task *model.Task) error {
	return u.taskRepo.Create(ctx, task)

}

func (u *taskUsecase) Update(ctx context.Context, task *model.Task) error {
	_, err := u.taskRepo.GetByID(ctx, task.ID)
	if err != nil {
		return errors.New("task not found")
	}
	return u.taskRepo.Update(ctx, task)
}

func (u *taskUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := u.taskRepo.GetByID(ctx, id)
	if err != nil {
		return errors.New("task not found")
	}
	return u.taskRepo.Delete(ctx, id)
}

func (u *taskUsecase) GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	return u.taskRepo.GetByID(ctx, id)
}

func (u *taskUsecase) GetAll(ctx context.Context, userID uuid.UUID, filter repository.TaskFilter) ([]*model.Task, error) {
	return u.taskRepo.GetAll(ctx, userID, filter)
}
