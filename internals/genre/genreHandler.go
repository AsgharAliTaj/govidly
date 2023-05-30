package genre

import (
	"encoding/json"
	"net/http"

	"github.com/asgharalitaj/govidly/database"
)

type GenreHandler struct {
	genreRepository GenreRepository
}

func NewGenreHandler() *GenreHandler {
	return &GenreHandler{
		genreRepository: NewGenreRepository(database.InitDB()),
	}
}

func (g *GenreHandler) GenreGetAll(w http.ResponseWriter, r *http.Request) {
	values, err := g.genreRepository.GetAllGenre()
	b, err := json.Marshal(values)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (g *GenreHandler) GenreGet(w http.ResponseWriter, r *http.Request) (err error) {
	// len := r.ContentLength
	//
	// body := make([]byte, len)
	// uuidString := string(body)
	//
	// r.Body.Read(body)
	// value, err := g.genreRepository.GetGenre(uuidString)
	// b, err := json.Marshal(value)
	// if err != nil {
	// 	return err
	// }
	// w.WriteHeader(200)
	// w.Write(b)
	// return nil
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
