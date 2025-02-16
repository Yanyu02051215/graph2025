package resolver

import (
	"context"
	"grapql-to-do/graph/model"
	"grapql-to-do/internal"
)

type TodoMutationResolver struct{}

func (r *TodoMutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	query := `INSERT INTO todos (text, done, user_id) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := internal.DB.QueryRow(ctx, query, input.Text, false, input.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:   int32(id),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID: input.UserID,
		},
	}, nil
}
