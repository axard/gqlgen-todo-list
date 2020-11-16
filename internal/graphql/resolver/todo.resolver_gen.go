package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/axard/gqlgen-todo-list/internal/db"
	"github.com/axard/gqlgen-todo-list/internal/graphql/dataloader"
	"github.com/axard/gqlgen-todo-list/internal/graphql/server"
	"github.com/axard/gqlgen-todo-list/internal/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return db.CreateTodo(ctx, input)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return db.Todos(ctx)
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return dataloader.For(ctx).UserByID.Load(obj.UserID)
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

// nolint: godox
// Todo returns server.TodoResolver implementation.
func (r *Resolver) Todo() server.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
