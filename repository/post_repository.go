package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

type PostRepositoryImpl struct {
}

func NewPostRepository() PostRepository {
	return &PostRepositoryImpl{}
}

func (repository *PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.Post, error) {
	query := "SELECT * FROM posts"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var posts []*model.Post
	for rows.Next() {
		post := model.Post{
			User: &model.User{},
		}
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.User.ID)
		if err != nil {
			return nil, err
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func (repository *PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (*model.Post, error) {
	query := "SELECT * FROM posts WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	post := model.Post{
		User: &model.User{},
	}
	err := row.Scan(&post.ID, &post.Title, &post.Body, &post.User.ID)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (repository *PostRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, post *model.Post) (*model.Post, error) {
	var id string
	query := "INSERT INTO posts (title, body, user_id) VALUES ($1, $2, $3) RETURNING id"
	row := tx.QueryRowContext(ctx, query, post.Title, post.Body, &post.User.ID)
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	post.ID = id

	return post, nil
}
