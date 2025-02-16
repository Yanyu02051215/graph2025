package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/domain/query"
	"grapql-to-do/graph/resolver/usecase"
)

type getTodosUsecase struct {
	todoQuery query.TodosQuery
}

func NewGetTodosUsecase(todoQuery query.TodosQuery) usecase.GetTodosUsecase {
	return &getTodosUsecase{todoQuery: todoQuery}
}

func (u *getTodosUsecase) GetTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := u.todoQuery.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
