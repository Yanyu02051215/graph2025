package usecase

import (
	"context"
	"grapql-to-do/internal/domain/model"
	"grapql-to-do/internal/domain/repository"

	model2 "grapql-to-do/graph/model"
)

type UserUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *model2.User) (*model.User, error) {
	modelUser := &model.User{Name: user.Name}
	id, err := u.userRepo.Create(ctx, modelUser)
	if err != nil {
		return nil, err
	}
	modelUser.ID = id
	return modelUser, nil
}
