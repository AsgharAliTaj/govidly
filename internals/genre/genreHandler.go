package genre

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/google/uuid"

	"github.com/asgharalitaj/govidly/database"
)

var (
	listGenres = regexp.MustCompile(`^\/genre[\/]*$`)   // /genre
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
	return nil
}

func (g *GenreHandler) GenreGet(w http.ResponseWriter, r *http.Request) (err error) {
	return nil
}

func (g *GenreHandler) GenrePost(w http.ResponseWriter, r *http.Request) (err error) {
	var genreName struct {
		Name string `json:"name"`
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &genreName)

	var genre Genre
	genre.ID = uuid.New()
	genre.Name = genreName.Name
	err = validateName(genre)
	if err != nil {
		return err
	}

	err = g.genreRepository.CreateGenre(genre)
	if err != nil {
		return err
	}
	return nil
}

func (g *GenreHandler) GenrePut(w http.ResponseWriter, r *http.Request) (err error) {
	return nil
}

func (g *GenreHandler) GenreDelete(w http.ResponseWriter, r *http.Request) (err error) {
	return nil
}
