package pg

import (
	"context"
	"os"

	"github.com/axard/gqlgen-todo-list/internal/log"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	defaultURL = "postgres://postgres:docker@localhost.ru:5432/postgres?sslmode=disable"
)

var (
	URL  string        = url()
	Pool *pgxpool.Pool = createConnectionPool()
)

func url() string {
	if v := os.Getenv("PG_URL"); v != "" {
		return v
	}

	return defaultURL
}

func createConnectionPool() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(), URL)
	if err != nil {
		log.Logger.Fatal(
			"PostgreSQL connect failed",
			zap.String("error", err.Error()),
			zap.String("URL", err.Error()),
		)
	}

	log.Logger.Info(
		"PostgreSQL connect success",
		zap.String("URL", err.Error()),
	)

	return pool
}
