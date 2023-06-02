package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
	"github.com/asgharalitaj/govidly/routes"
)

func main() {
	database.DatabseSetup()
	r := chi.NewRouter()
	routes.RegisterRoutes(r)

	server := http.Server{
		Addr:    ":3000",
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("server failed to start at port 3000")
	}
}
