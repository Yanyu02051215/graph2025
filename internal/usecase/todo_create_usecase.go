package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/domain/repository"

	model2 "grapql-to-do/graph/model"
)

type TodoCreateUsecase struct {
	todoRepo repository.TodoRepository
}

func NewCreateTodoUsecase(todoRepo repository.TodoRepository) *TodoCreateUsecase {
	return &TodoCreateUsecase{todoRepo: todoRepo}
}

func (u *TodoCreateUsecase) CreateTodo(ctx context.Context, todo *model2.Todo) (*model.Todo, error) {
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
