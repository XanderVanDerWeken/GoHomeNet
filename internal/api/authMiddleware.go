package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/xandervanderweken/GoHomeNet/internal/auth"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

type contextKey string

const UserContextKey = contextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			shared.WriteError(w, auth.ErrMissingHeader)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			shared.WriteError(w, auth.ErrWrongHeaderFormat)
			return
		}
		tokenStr := parts[1]

		claims, err := auth.ParseToken(tokenStr)
		if err != nil {
			shared.WriteError(w, auth.ErrInvalidToken)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
