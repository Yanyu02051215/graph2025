package container

import (
	"grapql-to-do/graph/resolver"
	"grapql-to-do/graph/schema"
	"grapql-to-do/internal/infrastructure/database"
	"grapql-to-do/internal/usecase"
)

// Resolver は全リゾルバを統合する
type Resolver struct {
	*resolver.UserMutationResolver
	*resolver.TodoMutationResolver
	*resolver.TodoQueryResolver
}

func NewResolver(db *database.Database) *Resolver {
	userDAO := database.NewUserCreateDAO(db)
	userUsecase := usecase.NewUserCreateUsecase(userDAO)
	userMutationResolver := resolver.NewUserMutationResolver(userUsecase)

	todoDAO := database.NewCreateTodoDAO(db)
	todoUsecase := usecase.NewCreateTodoUsecase(todoDAO)
	todoMutationResolver := resolver.NewTodoMutationResolver(todoUsecase)

	return &Resolver{
		UserMutationResolver: userMutationResolver,
		TodoMutationResolver: todoMutationResolver,
		TodoQueryResolver:    &resolver.TodoQueryResolver{},
	}
}

func (r *Resolver) Mutation() schema.MutationResolver {
	return r
}

func (r *Resolver) Query() schema.QueryResolver {
	return r
}
