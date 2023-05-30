package genre

import (
	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
)

func RegisterRoutes(r *chi.Mux) {
	genreHandler := &GenreHandler{
		genreRepository: NewGenreRepository(database.InitDB()),
	}

	r.Get("/genre", genreHandler.GenreGetAll)
}
