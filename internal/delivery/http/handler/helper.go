package handler

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/zhalisher/ip-task-manager/internal/middleware"
)

func getUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("userID not found in context")
	}
	return userID, nil
}
