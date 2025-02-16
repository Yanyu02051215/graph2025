package resolver

import (
	"context"

	"grapql-to-do/graph/model"
	"grapql-to-do/graph/resolver/usecase"
)

type UserMutationResolver struct {
	userUsecase usecase.CreateUserUsecase
}

func NewUserMutationResolver(userUsecase usecase.CreateUserUsecase) *UserMutationResolver {
	return &UserMutationResolver{userUsecase: userUsecase}
}

func (r *UserMutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user := &model.User{Name: input.Name}

	user2, err := r.userUsecase.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = user2.ID
	return user, nil
}
