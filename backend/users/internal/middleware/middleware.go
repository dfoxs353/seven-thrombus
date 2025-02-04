package middleware

import (
	"context"
	"log/slog"
	"main/internal/jwt"
	"main/internal/users"
	"net/http"
	"slices"
	"strings"

	"gitlab.com/volgaIt/packages/errorx"
	"gitlab.com/volgaIt/packages/middleware"
)

// доделать локальную штуку
func WrapAuth(
	handler http.Handler,
	validator jwt.TokenManager,
	requiredUserRole []users.Role,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(parts) != 2 {
			middleware.SendError(w, errorx.Unauthorized)
			return
		}

		if parts[0] != "Bearer" {
			middleware.SendError(w, errorx.Unauthorized)
			return
		}

		token := parts[1]

		payload, err := validator.ParseToken(token)
		if err != nil {
			slog.ErrorContext(r.Context(), "error while wrapping auth", "err", err)
			middleware.SendError(w, errorx.Unauthorized)
			return
		}

		var match bool
		for _, role := range requiredUserRole {
			if slices.Contains(payload.Roles, string(role)) {
				match = true
			}
		}
		if !match && len(requiredUserRole) != 0 {
			middleware.SendError(w, errorx.Forbidden)
			return
		}

		handler.ServeHTTP(w, r.WithContext(ContextWithToken(r.Context(), Token{payload.Uid, payload.Roles})))
	}
}

type Token struct {
	Uid   int
	Roles []string
}

type Key string

const token Key = "token"

func ContextWithToken(ctx context.Context, t Token) context.Context {
	return context.WithValue(ctx, token, t)
}

func TokenFromContext(ctx context.Context) (Token, bool) {
	token, ok := ctx.Value(token).(Token)
	return token, ok
}
