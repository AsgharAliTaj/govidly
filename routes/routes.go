package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/asgharalitaj/govidly/internals/auth"
	"github.com/asgharalitaj/govidly/internals/customer"
	"github.com/asgharalitaj/govidly/internals/genre"
	"github.com/asgharalitaj/govidly/internals/movie"
	"github.com/asgharalitaj/govidly/internals/user"
	"github.com/asgharalitaj/govidly/middlewares"
)

func RegisterRoutes(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middlewares.ResponseHeaderMiddleware)

	genre.RegisterRoutes(r)
	customer.RegisterRoutes(r)
	movie.RegisterRoutes(r)
	user.RegisterRoute(r)
	auth.RegisterRoutes(r)
}
