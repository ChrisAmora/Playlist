package handlers

import (
	"context"
	"net/http"

	"github.com/betopompolo/project_playlist_server/data"
)

type AuthToken struct {
	Name            string
	IsAuthenticated bool
}

func Auth(uc data.JWTRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c := r.Header.Get("authorization")
			if c == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, _ := uc.Verify(r.Context(), c)

			ctx := context.WithValue(r.Context(), "user", &AuthToken{Name: token.Username, IsAuthenticated: true})

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *AuthToken {
	raw, _ := ctx.Value("user").(*AuthToken)
	return raw
}
