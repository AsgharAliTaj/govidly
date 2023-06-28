package auth

import (
	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
)

func RegisterRoutes(r *chi.Mux) {
	authHandler := &AuthHandler{
		authRepository: &authRepository{
			database: database.InitDB(),
		},
	}

	r.Route("/api/auth", func(r chi.Router) {
		r.Get("/", authHandler.UserGet)
	})
}
