package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/samuelralmeida/neofarma/internal/auth"
)

type AuthMiddleware interface {
	SetUserToContext(ctx context.Context, userId string) (context.Context, error)
}

func SetUser(am AuthMiddleware) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token, err := auth.ReadCookie(r)
			if err != nil {
				log.Println("set user middleware reading cookie:", err)
				next.ServeHTTP(w, r)
				return
			}

			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			userID, err := auth.ParseToken(token)
			if err != nil {
				log.Println("set user middleware parsing token:", err)
				next.ServeHTTP(w, r)
				return
			}

			ctx, err := am.SetUserToContext(r.Context(), userID)
			if err != nil {
				log.Println("set user middleware parsing token:", err)
				next.ServeHTTP(w, r)
				return
			}

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
