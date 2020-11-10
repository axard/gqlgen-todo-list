package pg

import (
	"context"
	"time"

	"github.com/axard/gqlgen-todo-list/internal/cfg"
	"github.com/axard/gqlgen-todo-list/internal/log"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DefaultTimeout = 5 * time.Second
)

var (
	Pool *pgxpool.Pool
)

func init() {
	var err error

	config, err := pgxpool.ParseConfig(cfg.PgURL())
	if err != nil {
		log.Logger.Fatal(
			"PostgreSQL connect failed",
			zap.String("error", err.Error()),
			zap.String("URL", cfg.PgURL()),
		)
	}

	config.ConnConfig.Logger = zapadapter.NewLogger(log.Logger)

	Pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Logger.Fatal(
			"PostgreSQL connect failed",
			zap.String("error", err.Error()),
			zap.String("URL", cfg.PgURL()),
		)
	}

	log.Logger.Info(
		"PostgreSQL connect success",
		zap.String("URL", cfg.PgURL()),
	)
}

func Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, DefaultTimeout)
	defer cancelFunc()

	conn, err := Pool.Acquire(ctxWithTimeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
