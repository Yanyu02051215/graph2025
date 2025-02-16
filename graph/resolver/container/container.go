package container

import (
	"grapql-to-do/graph/resolver"
	"grapql-to-do/graph/schema"
	"grapql-to-do/internal/infrastructure"
)

// Resolver は全リゾルバを統合する
type Resolver struct {
	*resolver.UserMutationResolver
	*resolver.TodoMutationResolver
	*resolver.TodoQueryResolver
}

// func NewResolver(db *infrastructure.Database) *Resolver {
// 	userDAO := database.NewUserCreateDAO(db)
// 	userUsecase := usecase.NewUserCreateUsecase(userDAO)
// 	userMutationResolver := resolver.NewUserMutationResolver(userUsecase)

// 	createTodoDAO := database.NewCreateTodoDAO(db)
// 	createTodoUsecase := usecase.NewCreateTodoUsecase(createTodoDAO)
// 	todoMutationResolver := resolver.NewTodoMutationResolver(createTodoUsecase)

// 	getTodoDAO := database.NewGetTodoDAO(db)
// 	getTodoUsecase := usecase.NewGetTodosUsecase(getTodoDAO)
// 	todoQueryResolver := resolver.NewTodoQueryResolver(getTodoUsecase)

// 	return &Resolver{
// 		UserMutationResolver: userMutationResolver,
// 		TodoMutationResolver: todoMutationResolver,
// 		TodoQueryResolver:    todoQueryResolver,
// 	}
// }

func NewResolver(db *infrastructure.Database) *Resolver {
	return &Resolver{
		UserMutationResolver: newUserMutationResolver(db),
		TodoMutationResolver: newTodoMutationResolver(db),
		TodoQueryResolver:    newTodoQueryResolver(db),
	}
}

func (r *Resolver) Mutation() schema.MutationResolver {
	return r
}

func (r *Resolver) Query() schema.QueryResolver {
	return r
}
