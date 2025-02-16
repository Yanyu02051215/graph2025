package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/domain/query"
)

type GetTodosUsecase struct {
	todoQuery query.TodosQuery
}

func NewGetTodosUsecase(todoQuery query.TodosQuery) *GetTodosUsecase {
	return &GetTodosUsecase{todoQuery: todoQuery}
}

func (u *GetTodosUsecase) GetTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := u.todoQuery.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
