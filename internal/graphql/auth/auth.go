package auth

import (
	"context"
	"net/http"
	"strconv"

	"github.com/axard/gqlgen-todo-list/internal/log"
	"go.uber.org/zap"
)

type ContextKey string

const key ContextKey = "authorization"

type Payload struct {
	UserID int
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.Header.Get("Authorization")
		if value == "" {
			log.Logger.Error(
				"invalid Authorization header",
			)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}

		userID, err := strconv.Atoi(value)
		if err != nil {
			log.Logger.Error(
				"invalid Authorization header",
				zap.String("error", err.Error()),
			)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)

			return
		}

		ctx := context.WithValue(r.Context(), key, &Payload{
			UserID: userID,
		})

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Payload {
	return ctx.Value(key).(*Payload)
}
