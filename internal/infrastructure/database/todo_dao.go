package database

import (
	"context"
	"grapql-to-do/internal/domain/model"
)

type TodoDAO struct {
	DB *Database
}

func NewCreateTodoDAO(db *Database) *TodoDAO {
	return &TodoDAO{DB: db}
}

func (dao *TodoDAO) Create(ctx context.Context, todo *model.Todo) (*model.Todo, error) {
	query := `INSERT INTO todos (text, done, user_id) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := dao.DB.Conn.QueryRow(ctx, query, todo.Text, false, todo.UserId).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:     int32(id),
		Text:   todo.Text,
		Done:   false,
		UserId: todo.UserId,
	}, nil
}
