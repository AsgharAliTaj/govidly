package genre

import (
	"net/http"
	"regexp"

	"github.com/asgharalitaj/govidly/database"
)

var (
	listGenres = regexp.MustCompile(`^\/genres[\/]*$`)  // /genre
	getGenre   = regexp.MustCompile(`^\/genre\/{\d+}$`) // /genre/123
)

type GenreHandler struct {
	genreRepository GenreRepository
}

func NewGenreHandler() *GenreHandler {
	return &GenreHandler{
		genreRepository: NewGenreRepository(database.InitDB()),
	}
}

func (g *GenreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch {
	case r.Method == http.MethodGet && listGenres.MatchString(r.URL.Path):
		err = g.GenreGetAll(w, r)
	case r.Method == http.MethodGet && getGenre.MatchString(r.URL.Path):
		err = g.GenreGet(w, r)
	case r.Method == http.MethodPost:
		err = g.GenrePost(w, r)
	case r.Method == http.MethodPut:
		err = g.GenrePut(w, r)
	case r.Method == http.MethodDelete:
		err = g.GenreDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (g *GenreHandler) GenreGetAll(w http.ResponseWriter, r *http.Request) (err error) {
}

func (g *GenreHandler) GenreGet(w http.ResponseWriter, r *http.Request) (err error) {
}

func (g *GenreHandler) GenrePost(w http.ResponseWriter, r *http.Request) (err error) {
}

func (g *GenreHandler) GenrePut(w http.ResponseWriter, r *http.Request) (err error) {
}

func (g *GenreHandler) GenreDelete(w http.ResponseWriter, r *http.Request) (err error) {
}
