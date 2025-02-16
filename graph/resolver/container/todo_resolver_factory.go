package container

import (
	"grapql-to-do/graph/resolver"
	"grapql-to-do/internal/infrastructure"
	"grapql-to-do/internal/infrastructure/database"
	"grapql-to-do/internal/usecase"
)

func newTodoMutationResolver(db *infrastructure.Database) *resolver.TodoMutationResolver {
	createTodoDAO := database.NewCreateTodoDAO(db)
	createTodoUsecase := usecase.NewCreateTodoUsecase(createTodoDAO)
	return resolver.NewTodoMutationResolver(createTodoUsecase)
}

func newTodoQueryResolver(db *infrastructure.Database) *resolver.TodoQueryResolver {
	getTodoDAO := database.NewGetTodoDAO(db)
	getTodoUsecase := usecase.NewGetTodosUsecase(getTodoDAO)
	return resolver.NewTodoQueryResolver(getTodoUsecase)
}
