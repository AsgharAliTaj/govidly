package genre

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GenreHandler struct {
	genreRepository Genrer
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
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	var genreName struct {
		Name string `json:"name"`
	}

	var genre Genre

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &genreName)
	genre.Name = genreName.Name

	err := g.genreRepository.CreateGenre(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (g *GenreHandler) GenrePut(w http.ResponseWriter, r *http.Request) {
	// not implemented because i don't need this for this small api of genre
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
