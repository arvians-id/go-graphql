package resolver

import (
	"context"
	"github.com/arvians-id/go-graphql/controller/graph/generated"
	"github.com/arvians-id/go-graphql/controller/graph/model"
)

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]*model.Post, error) {
	postsResolver, err := r.PostService.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var posts []*model.Post
	for _, post := range postsResolver {
		if post.User.ID == obj.ID {
			posts = append(posts, post)
		}
	}

	return posts, nil
}
