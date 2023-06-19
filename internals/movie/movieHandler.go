package movie

import (
	"encoding/json"
	"net/http"
)

type MovieHandler struct {
	MovieRepository Movier
}

func (m *MovieHandler) MovieGetAll(w http.ResponseWriter, r *http.Request) {
	movies, err := m.MovieRepository.GetAllMovies()
	b, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
