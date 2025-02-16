package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	model2 "grapql-to-do/graph/model"
)

type CreateUserUsecase interface {
	CreateUser(ctx context.Context, todo *model2.User) (*model.User, error)
}
