package dataloader

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/axard/gqlgen-todo-list/internal/model"
)

type ContextKey string

const key ContextKey = "dataloader"

const (
	DefaultBatchSize = 100
	DefaultWaitTime  = 1 * time.Millisecond
)

type Payload struct {
	UserByID UserByID
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), key, &Payload{
			UserByID: UserByID{
				maxBatch: DefaultBatchSize,
				wait:     DefaultWaitTime,
				fetch: func(keys []int) ([]*model.User, []error) {
					panic(fmt.Errorf("not implemented"))
				},
			},
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Payload {
	return ctx.Value(key).(*Payload)
}
