package genre

import (
	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
	"github.com/asgharalitaj/govidly/middlewares"
)

func RegisterRoutes(r *chi.Mux) {
	genreHandler := &GenreHandler{
		genreRepository: &genreRepository{
			database: database.InitDB(),
		},
	}

	r.Route("/api/genres", func(r chi.Router) {
		r.Get("/", genreHandler.GenreGetAll)
		r.Get("/{genreId}", genreHandler.GenreGet)
		r.Post("/", middlewares.AuthStatusCheckingMiddleware(genreHandler.GenrePost))
		r.Delete("/{genreId}", genreHandler.GenreDelete)
	})
}
