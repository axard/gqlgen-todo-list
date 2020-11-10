package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/axard/gqlgen-todo-list/internal/graphql/generated"
	"github.com/axard/gqlgen-todo-list/internal/graphql/model"
	"github.com/axard/gqlgen-todo-list/internal/pg"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	conn, err := pg.Acquire(ctx)
	if err != nil {
		return nil, errors.New("Unable to handle request")
	}

	const sql = `INSERT INTO todos (description, user_id) VALUES ($1, $2) RETURNING id, done`

	rows, err := conn.Query(ctx, sql, input.Description, input.UserID)
	if err != nil {
		return nil, errors.New("Unable to handle request")
	}

	defer rows.Close()

	var (
		id   int
		done bool
	)

	for rows.Next() {
		if err := rows.Scan(&id, &done); err != nil {
			return nil, errors.New("Unable to handle request")
		}
	}

	return &model.Todo{
		ID:          id,
		Description: input.Description,
		Done:        done,
		UserID:      input.UserID,
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	conn, err := pg.Acquire(ctx)
	if err != nil {
		return nil, errors.New("Unable to handle request")
	}

	const sql = `SELECT id, description, done, user_id FROM todos`

	rows, err := conn.Query(ctx, sql)
	if err != nil {
		return nil, errors.New("Unable to handle request")
	}

	defer rows.Close()

	var todos []*model.Todo

	for rows.Next() {
		todo := &model.Todo{}

		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Done, &todo.UserID); err != nil {
			return nil, errors.New("Unable to handle request")
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	conn, err := pg.Acquire(ctx)
	if err != nil {
		return nil, errors.New("Unable to handle request")
	}

	const sql = `SELECT login FROM users WHERE id = $1`

	rows, err := conn.Query(ctx, sql, obj.UserID)
	if err != nil {
		return nil, errors.New("Unable to handle request")
	}

	defer rows.Close()

	var login string

	for rows.Next() {
		if err := rows.Scan(&login); err != nil {
			return nil, errors.New("Unable to handle request")
		}
	}

	return &model.User{
		ID:    obj.UserID,
		Login: login,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// nolint: godox
// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
