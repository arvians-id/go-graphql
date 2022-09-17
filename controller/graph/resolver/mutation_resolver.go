package resolver

import (
	"context"
	"errors"
	"github.com/arvians-id/go-graphql/controller/graph/generated"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name: input.Name,
	}

	user, err := r.UserService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	user, err := r.UserService.FindById(ctx, input.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	post := &model.Post{
		Title: input.Title,
		Body:  input.Body,
		User:  user,
	}

	post, err = r.PostService.Create(ctx, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
