package pg

import (
	"context"

	"github.com/axard/gqlgen-todo-list/internal/cfg"
	"github.com/axard/gqlgen-todo-list/internal/log"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	Pool *pgxpool.Pool
)

func init() {
	var err error

	Pool, err = pgxpool.Connect(context.Background(), cfg.PgURL())
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
}
