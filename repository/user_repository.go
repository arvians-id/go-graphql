package repository

import (
	"context"
	"database/sql"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*model.User, error) {
	query := "SELECT * FROM users"
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

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id string) (*model.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := tx.QueryRowContext(ctx, query, id)

	var user model.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user *model.User) (*model.User, error) {
	var id string
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	row := tx.QueryRowContext(ctx, query, user.Name)
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return user, nil
}
