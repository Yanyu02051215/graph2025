// package resolver

// import (
// 	"context"
// 	"grapql-to-do/graph/model"
// 	"grapql-to-do/internal"
// )

// type TodoQueryResolver struct{}

// func (r *TodoQueryResolver) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
// 	query := `SELECT
// 				todos.id,
// 				todos.text,
// 				todos.done,
// 				todos.user_id,
// 				users.name
// 			FROM todos
// 			INNER JOIN users ON users.id = todos.user_id`
// 	rows, err := internal.DB.Query(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var todos []*model.Todo
// 	for rows.Next() {
// 		var todo model.Todo
// 		var userID int32
// 		var userName string
// 		err := rows.Scan(&todo.ID, &todo.Text, &todo.Done, &userID, &userName)
// 		if err != nil {
// 			return nil, err
// 		}

// 		todo.User = &model.User{
// 			ID:   userID,
// 			Name: userName,
// 		}
// 		todos = append(todos, &todo)
// 	}

// 	return todos, nil
// }

package resolver

import (
	"context"

	"grapql-to-do/graph/model"
	"grapql-to-do/internal/usecase"
)

type TodoQueryResolver struct {
	todoUsecase *usecase.GetTodosUsecase
}

func NewTodoQueryResolver(todoUsecase *usecase.GetTodosUsecase) *TodoQueryResolver {
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
