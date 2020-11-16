package main

import (
	"net/http"
	"os"

	"github.com/axard/gqlgen-todo-list/internal/cfg"
	"github.com/axard/gqlgen-todo-list/internal/db/pg"
	"github.com/axard/gqlgen-todo-list/internal/graphql/dataloader"
	"github.com/axard/gqlgen-todo-list/internal/graphql/resolver"
	"github.com/axard/gqlgen-todo-list/internal/graphql/server"
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

	if err := pg.Connect(cfg.PgURL()); err != nil {
		log.Logger.Error(
			"PostgreSQL connect failed",
			zap.String("error", err.Error()),
			zap.String("URL", cfg.PgURL()),
		)
	}

	log.Logger.Info(
		"PostgreSQL connect success",
		zap.String("URL", cfg.PgURL()),
	)

	srv := handler.NewDefaultServer(server.NewExecutableSchema(
		server.Config{
			Resolvers: &resolver.Resolver{},
		},
	))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.Middleware(srv))

	log.Logger.Sugar().Infof(
		"connect to http://localhost:%s/ for GraphQL playground", port)
	log.Logger.Fatal(http.ListenAndServe(":"+port, nil).Error())
}
