package main

import (
	"net/http"

	"github.com/asgharalitaj/govidly/database"
)

func main() {
	database.DatabseSetup()
	mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	server.ListenAndServe()
}
