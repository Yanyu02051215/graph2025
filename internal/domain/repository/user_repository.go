package repository

import (
	"context"
	"grapql-to-do/internal/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int32, error)
}
