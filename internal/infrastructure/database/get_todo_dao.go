package database

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/infrastructure"
)

type GetTodoDAO struct {
	DB *infrastructure.Database
}

func NewGetTodoDAO(db *infrastructure.Database) *GetTodoDAO {
	return &GetTodoDAO{DB: db}
}

func (dao *GetTodoDAO) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	query := `SELECT
				todos.id,
				todos.text,
				todos.done,
				todos.user_id,
				users.name
			FROM todos
			INNER JOIN users ON users.id = todos.user_id`
	rows, err := dao.DB.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []*model.Todo
	for rows.Next() {
		var todo model.Todo
		var userID int32
		var userName string
		err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &userID, &userName)
		if err != nil {
			return nil, err
		}

		todo.User = &model.User{
			ID:   userID,
			Name: userName,
		}
		todos = append(todos, &todo)
	}

	return todos, nil
}
