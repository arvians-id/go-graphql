package main

import (
	"database/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arvians-id/go-graphql/config"
	"github.com/arvians-id/go-graphql/controller/graph/generated"
	"github.com/arvians-id/go-graphql/controller/graph/resolver"
	"github.com/arvians-id/go-graphql/repository"
	"github.com/arvians-id/go-graphql/service"
	"github.com/gin-gonic/gin"
	"log"
)

func NewInitializedDatabase() (*sql.DB, error) {
	db, err := config.NewPostgreSQL()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func graphqlHandler() gin.HandlerFunc {
	db, _ := NewInitializedDatabase()
	userRepository := repository.NewUserRepository()
	postRepository := repository.NewPostRepository()

	userService := service.NewUserService(&userRepository, db)
	postService := service.NewPostService(&postRepository, db)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{
			UserService: userService,
			PostService: postService,
		},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// Setting up Gin
	r := gin.Default()
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	err := r.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
