package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/domain/repository"
	"grapql-to-do/graph/resolver/usecase"

	model2 "grapql-to-do/graph/model"
)

type createUserUsecaseImpl struct {
	userRepo repository.UserRepository
}

func NewUserCreateUsecase(userRepo repository.UserRepository) usecase.CreateUserUsecase {
	return &createUserUsecaseImpl{userRepo: userRepo}
}

func (u *createUserUsecaseImpl) CreateUser(ctx context.Context, user *model2.User) (*model.User, error) {
	modelUser := &model.User{Name: user.Name}
	id, err := u.userRepo.Create(ctx, modelUser)
	if err != nil {
		return nil, err
	}
	modelUser.ID = id
	return modelUser, nil
}
