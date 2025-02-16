package database

import (
	"context"
	"grapql-to-do/internal/domain/model"
)

type CreateTodoDAO struct {
	DB *Database
}

func NewCreateTodoDAO(db *Database) *CreateTodoDAO {
	return &CreateTodoDAO{DB: db}
}

func (dao *CreateTodoDAO) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	query := `INSERT INTO todos (text, done, user_id) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := dao.DB.Conn.QueryRow(ctx, query, todo.Text, false, todo.User.ID).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:     int32(id),
		Text:   todo.Text,
		Done:   false,
		User: &model.User{
			ID:   todo.User.ID,
		},
	}, nil
}
