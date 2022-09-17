package resolver

import (
	"context"
	"errors"
	"github.com/arvians-id/go-graphql/controller/graph/generated"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.UserService.FindById(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts, err := r.PostService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	post, err := r.PostService.FindById(ctx, id)
	if err != nil {
		return nil, errors.New("post not found")
	}

	return post, nil
}
