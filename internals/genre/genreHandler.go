package genre

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (g *GenreHandler) GenreGet(w http.ResponseWriter, r *http.Request) {
	genreId := chi.URLParam(r, "genreId")
	value, err := strconv.Atoi(genreId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	genre, err := g.genreRepository.GetGenre(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonEncodedGenre, err := json.Marshal(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonEncodedGenre)
}

func (g *GenreHandler) GenrePost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GenrePost handler")

	var genreName struct {
		Name string `json:"name"`
	}

	var genre Genre

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &genreName)
	genre.Name = genreName.Name

	// err := validateName(genre)
	// if err != nil {
	// 	fmt.Println(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	err := g.genreRepository.CreateGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (g *GenreHandler) GenrePut(w http.ResponseWriter, r *http.Request) {
}

func (g *GenreHandler) GenreDelete(w http.ResponseWriter, r *http.Request) {
	genreId := chi.URLParam(r, "genreId")
	value, err := strconv.Atoi(genreId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = g.genreRepository.DeleteGenre(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
