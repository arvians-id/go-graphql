package service

import (
	"context"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

type UserService interface {
	FindAll(ctx context.Context) ([]*model.User, error)
	FindById(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
}

type PostService interface {
	FindAll(ctx context.Context) ([]*model.Post, error)
	FindById(ctx context.Context, id string) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (*model.Post, error)
}
