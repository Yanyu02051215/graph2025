package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	model2 "grapql-to-do/graph/model"
)

type TodoUsecase interface {
	CreateTodo(ctx context.Context, todo *model2.Todo) (*model.Todo, error)
}

type GetTodosUsecase interface {
	GetTodos(ctx context.Context) ([]*model.Todo, error)
}
