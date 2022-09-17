//go:generate go run github.com/99designs/gqlgen generate
package resolver

import "github.com/arvians-id/go-graphql/service"

type Resolver struct {
	UserService service.UserService
	PostService service.PostService
}
