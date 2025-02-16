package resolver

import (
	"context"

	"grapql-to-do/graph/model"
	"grapql-to-do/graph/resolver/usecase"
)

type TodoMutationResolver struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoMutationResolver(todoUsecase usecase.TodoUsecase) *TodoMutationResolver {
	return &TodoMutationResolver{todoUsecase: todoUsecase}
}

func (r *TodoMutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		User: &model.User{
			ID: input.UserID,
		},
	}

	createdTodo, err := r.todoUsecase.CreateTodo(ctx, todo)
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:   createdTodo.ID,
		Text: createdTodo.Text,
		Done: createdTodo.Done,
		User: &model.User{ID: createdTodo.User.ID},
	}, nil
}
