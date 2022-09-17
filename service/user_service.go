package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-graphql/controller/graph/model"
	"github.com/arvians-id/go-graphql/helper"
	"github.com/arvians-id/go-graphql/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
}

func NewUserService(userRepository *repository.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepository: *userRepository,
		DB:             db,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	users, err := service.UserRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id string) (*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service *UserServiceImpl) Create(ctx context.Context, user *model.User) (*model.User, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	user, err = service.UserRepository.Create(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
