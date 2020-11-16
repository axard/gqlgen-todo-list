package db

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/axard/gqlgen-todo-list/internal/db/pg"
	"github.com/axard/gqlgen-todo-list/internal/log"
	"github.com/axard/gqlgen-todo-list/internal/model"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

var (
	ErrCreateFail = errors.New("create fail")
	ErrFetchFail  = errors.New("fetch fail")
)

func CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	sql, args, err := pg.SQL.
		Insert("todos").
		Columns("description", "user_id").
		Values(input.Description, input.UserID).
		Suffix("RETURNING id, done").
		ToSql()
	if err != nil {
		log.Logger.Error(
			"on build sql string",
			zap.String("error", err.Error()),
		)

		return nil, ErrCreateFail
	}

	var todo model.Todo

	err = pg.NewQuery(sql).SetArgs(args).SetScan(&todo, pgxscan.ScanOne).Execute(ctx)
	if err != nil {
		log.Logger.Error(
			"on execute sql query",
			zap.String("error", err.Error()),
		)

		return nil, ErrCreateFail
	}

	todo.Description = input.Description
	todo.UserID = input.UserID

	return &todo, nil
}

func Todos(ctx context.Context) ([]*model.Todo, error) {
	sql, _, err := pg.SQL.
		Select("id", "description", "done", "user_id").
		From("todos").
		ToSql()
	if err != nil {
		log.Logger.Error(
			"on build sql string",
			zap.String("error", err.Error()),
		)

		return nil, ErrFetchFail
	}

	var todos []*model.Todo

	err = pg.NewQuery(sql).SetScan(&todos, pgxscan.ScanAll).Execute(ctx)
	if err != nil {
		log.Logger.Error(
			"on execute sql query",
			zap.String("error", err.Error()),
		)

		return nil, ErrCreateFail
	}

	return todos, nil
}

func UserById(ctx context.Context, id int) (*model.User, error) {
	sql, args, err := pg.SQL.
		Select("login").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Logger.Error(
			"on build sql string",
			zap.String("error", err.Error()),
		)

		return nil, ErrFetchFail
	}

	var user model.User

	err = pg.NewQuery(sql).SetArgs(args).SetScan(&user, pgxscan.ScanOne).Execute(ctx)
	if err != nil {
		log.Logger.Error(
			"on execute sql query",
			zap.String("error", err.Error()),
		)

		return nil, ErrCreateFail
	}

	user.ID = id

	return &user, nil
}

func UsersByIds(ctx context.Context, ids []int) ([]*model.User, error) {
	sql, args, err := pg.SQL.
		Select("id, login").
		From("users").
		Where(squirrel.Eq{"id": ids}).
		ToSql()
	if err != nil {
		log.Logger.Error(
			"on build sql string",
			zap.String("error", err.Error()),
		)

		return nil, ErrFetchFail
	}

	usersMap := make(map[int]*model.User)
	scanFunc := func(dest interface{}, rows pgx.Rows) error {
		for rows.Next() {
			var user model.User

			if err := pgxscan.ScanRow(&user, rows); err != nil {
				return err
			}

			usersMap[user.ID] = &user
		}

		return nil
	}

	err = pg.NewQuery(sql).SetArgs(args).SetScan(&usersMap, scanFunc).Execute(ctx)
	if err != nil {
		log.Logger.Error(
			"on execute sql query",
			zap.String("error", err.Error()),
		)

		return nil, ErrCreateFail
	}

	// Сортируем полученных пользователей в порядке полученных айдишников
	users := make([]*model.User, len(usersMap))
	for i, id := range ids {
		users[i] = usersMap[id]
	}

	return users, nil
}
