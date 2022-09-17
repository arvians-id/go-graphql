package resolver

import (
	"context"
	"github.com/arvians-id/go-graphql/controller/graph/generated"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	user, err := r.Resolver.UserService.FindById(ctx, obj.User.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type postResolver struct{ *Resolver }
