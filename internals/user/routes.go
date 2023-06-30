package user

import (
	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
	"github.com/asgharalitaj/govidly/middlewares"
)

func RegisterRoute(r *chi.Mux) {
	userHandler := &UserHandler{
		userRepository: &userRepository{
			database: database.InitDB(),
		},
	}

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", userHandler.UserPost)
		r.Get("/me", middlewares.AuthStatusCheckingMiddleware(userHandler.UserGet))
	})
}
