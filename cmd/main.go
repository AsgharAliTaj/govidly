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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to listen and serve ", err)
	}
	log.Println("server is up and running on Port:3000")
}
