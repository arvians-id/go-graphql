package service

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-graphql/controller/graph/model"
	"github.com/arvians-id/go-graphql/helper"
	"github.com/arvians-id/go-graphql/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
}

func NewPostService(postRepository *repository.PostRepository, db *sql.DB) PostService {
	return &PostServiceImpl{
		PostRepository: *postRepository,
		DB:             db,
	}
}

func (service *PostServiceImpl) FindAll(ctx context.Context) ([]*model.Post, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	posts, err := service.PostRepository.FindAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (service *PostServiceImpl) FindById(ctx context.Context, id string) (*model.Post, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (service *PostServiceImpl) Create(ctx context.Context, post *model.Post) (*model.Post, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer helper.CommitOrRollback(tx)

	post, err = service.PostRepository.Create(ctx, tx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
