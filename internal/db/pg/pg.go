package pg

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/axard/gqlgen-todo-list/internal/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	DefaultTimeout = 5 * time.Second
)

var (
	pool *pgxpool.Pool

	SQL = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func Connect(url string) error {
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return err
	}

	config.ConnConfig.Logger = zapadapter.NewLogger(log.Logger)

	pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}

	return nil
}

type ScanFunc func(interface{}, pgx.Rows) error

type Query struct {
	sql string

	scan ScanFunc

	args []interface{}
	dest interface{}
}

func NewQuery(sql string) *Query {
	return &Query{
		sql: sql,
	}
}

func (this *Query) SetArgs(args []interface{}) *Query {
	this.args = args
	return this
}

func (this *Query) SetScan(dest interface{}, scan ScanFunc) *Query {
	this.dest = dest
	this.scan = scan

	return this
}

func (this *Query) Execute(ctx context.Context) error {
	ctxWithTimeout, cancelFunc := context.WithTimeout(ctx, DefaultTimeout)
	defer cancelFunc()

	rows, err := pool.Query(ctxWithTimeout, this.sql, this.args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	if this.dest != nil && this.scan != nil {
		err := this.scan(this.dest, rows)
		if err != nil {
			return err
		}
	}

	return nil
}
