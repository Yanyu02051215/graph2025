package repository

import (
	"context"
	"grapql-to-do/internal/domain/model"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *model.Todo) (*model.Todo, error)
}
