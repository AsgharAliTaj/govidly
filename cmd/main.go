package main

import (
	"net/http"

	"github.com/asgharalitaj/govidly/internals/genre"
)

func main() {
	// database.DatabseSetup()
	genreHandler := genre.NewGenreHandler()

	mux := http.NewServeMux()
	mux.Handle("/api/genre", genreHandler)

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	server.ListenAndServe()
}
