package resolver

import (
	"context"

	"grapql-to-do/graph/model"
	"grapql-to-do/graph/resolver/usecase"
)

type TodoQueryResolver struct {
	todoUsecase usecase.GetTodosUsecase
}

func NewTodoQueryResolver(todoUsecase usecase.GetTodosUsecase) *TodoQueryResolver {
	return &TodoQueryResolver{todoUsecase: todoUsecase}
}

func (r *TodoQueryResolver) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.todoUsecase.GetTodos(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*model.Todo, len(todos))
	for i, todo := range todos {
		result[i] = &model.Todo{
			ID:   todo.ID,
			Text: todo.Text,
			Done: todo.Done,
			User: &model.User{
				ID:   todo.User.ID,
				Name: todo.User.Name,
			},
		}
	}

	return result, nil
}
