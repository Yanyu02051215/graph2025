package resolver

import (
	"context"

	"grapql-to-do/graph/model"
	"grapql-to-do/internal/usecase"
)

// UserMutationResolver は User のミューテーションリゾルバ
type UserMutationResolver struct {
	userUsecase *usecase.UserUsecase
}

// NewUserMutationResolver は UserMutationResolver を生成する
func NewUserMutationResolver(userUsecase *usecase.UserUsecase) *UserMutationResolver {
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
