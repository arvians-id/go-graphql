package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.User, error)
	FindById(ctx context.Context, tx *sql.Tx, id string) (*model.User, error)
	Create(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error)
}

type PostRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Post, error)
	FindById(ctx context.Context, tx *sql.Tx, id string) (*model.Post, error)
	Create(ctx context.Context, tx *sql.Tx, post *model.Post) (*model.Post, error)
}
