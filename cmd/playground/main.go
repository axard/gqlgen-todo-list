package main

import (
	"net/http"
	"os"

	"github.com/axard/gqlgen-todo-list/internal/graphql"
	"github.com/axard/gqlgen-todo-list/internal/graphql/generated"
	"github.com/axard/gqlgen-todo-list/internal/log"
	"github.com/axard/gqlgen-todo-list/pkg/version"
	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	log.Logger.Info(
		"Playground started",
		zap.String("Version", version.Version),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graphql.Resolver{},
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Logger.Sugar().Infof(
		"connect to http://localhost:%s/ for GraphQL playground", port)
	log.Logger.Fatal(http.ListenAndServe(":"+port, nil).Error())
}
