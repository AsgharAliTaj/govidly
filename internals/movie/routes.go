package movie

import (
	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
)

func RegisterRoutes(r *chi.Mux) {
	moviehandler := &MovieHandler{
		MovieRepository: &movieRepository{
			databse: database.InitDB(),
		},
	}
	r.Route("/api/movies", func(r chi.Router) {
		r.Get("/", moviehandler.MovieGetAll)
	})
}
