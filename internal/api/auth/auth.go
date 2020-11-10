package auth

import (
	"context"
	"net/http"

	"github.com/axard/gqlgen-todo-list/internal/api"
)

// TODO Сделать нормальную авторизацию
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth := r.Header.Get("Authorization")
		if auth != "" {
			ctx = context.WithValue(ctx, api.UserID, auth)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
