package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ResponseHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthStatusCheckingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Values("x-auth-token")
		if len(token) == 0 {
			http.Error(w, "access denied no token provided", http.StatusUnauthorized)
			return
		}
		var claims struct {
			Id uuid.UUID `json:"id"`
			jwt.RegisteredClaims
		}
		tkn, err := jwt.ParseWithClaims(
			token[0],
			&claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte("mysecretkey"), nil
			})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !tkn.Valid {
			http.Error(w, "Token is invalid", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "id", claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
