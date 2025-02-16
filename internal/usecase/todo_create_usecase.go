package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/domain/repository"
	"grapql-to-do/graph/resolver/usecase"

	model2 "grapql-to-do/graph/model"
)

type todoCreateUsecaseImpl struct {
	todoRepo repository.TodoRepository
}

func NewCreateTodoUsecase(todoRepo repository.TodoRepository) usecase.TodoUsecase {
	return &todoCreateUsecaseImpl{todoRepo: todoRepo}
}

func (u *todoCreateUsecaseImpl) CreateTodo(ctx context.Context, todo *model2.Todo) (*model.Todo, error) {
	modelTodo := &model.Todo{
		Text:   todo.Text,
		User: &model.User{
			ID:   todo.User.ID,
		},
	}
	todo2, err := u.todoRepo.Create(ctx, modelTodo)
	if err != nil {
		return nil, err
	}
	return todo2, nil
}
