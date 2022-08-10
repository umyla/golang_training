package middleware

import (
	"context"
	"errors"
	"net/http"
	"service/auth"
	"service/web"
	"strings"
)

func (m *Mid) Authenticate(next web.HandlerFunc) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		authHeader := r.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := errors.New("expexted authorization header format:bearer <token>")
			return web.NewRequestError(err, http.StatusUnauthorized)
		}
		claims, err := m.A.ValidateToken(parts[1])
		if err != nil {
			return web.NewRequestError(err, http.StatusUnauthorized)
		}
		ctx = context.WithValue(ctx, auth.Key, claims)
		return next(ctx, w, r)

	}
}
func (m *Mid) HasRole(next web.HandlerFunc, roles ...string) web.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		claims, ok := ctx.Value(auth.Key).(auth.Claims)
		if !ok {
			return errors.New("claims missing from context: HasRole called without/before Authenticate")
		}

		ok = claims.HasRoles(roles...)
		if !ok {
			return web.NewRequestError(errors.New("you are not authorized for that action"), http.StatusForbidden)
		}

		return next(ctx, w, r)

	}

}
