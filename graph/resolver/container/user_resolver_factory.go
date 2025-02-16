package container

import (
	"grapql-to-do/graph/resolver"
	"grapql-to-do/internal/infrastructure"
	"grapql-to-do/internal/infrastructure/database"
	"grapql-to-do/internal/usecase"
)

func newUserMutationResolver(db *infrastructure.Database) *resolver.UserMutationResolver {
	userDAO := database.NewUserCreateDAO(db)
	userUsecase := usecase.NewUserCreateUsecase(userDAO)
	return resolver.NewUserMutationResolver(userUsecase)
}
