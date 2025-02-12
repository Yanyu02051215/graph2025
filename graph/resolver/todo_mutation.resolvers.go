package resolver

import (
	"context"
	"grapql-to-do/graph/model"
	"grapql-to-do/internal"
)

// UserMutationResolver は User のミューテーションリゾルバ
type UserMutationResolver struct{}

// CreateUser は createUser ミューテーションのリゾルバ
func (r *UserMutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	query := `INSERT INTO users (name) VALUES ($1) RETURNING id`
	var id int
	err := internal.DB.QueryRow(ctx, query, input.Name).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:   int32(id),
		Name: input.Name,
	}, nil
}
