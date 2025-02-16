package query

import (
	"context"
	"grapql-to-do/internal/domain/model"
)

type TodosQuery interface {
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
}
