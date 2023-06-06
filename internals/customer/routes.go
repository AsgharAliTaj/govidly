package customer

import (
	"github.com/go-chi/chi/v5"

	"github.com/asgharalitaj/govidly/database"
)

func RegisterRoutes(r *chi.Mux) {
	customerHandler := &CustomerHandler{
		CustomerRepository: &CustomerRepository{
			database: database.InitDB(),
		},
	}

	r.Route("/api/customers", func(r chi.Router) {
		r.Get("/", customerHandler.CustomerGetAll)
		r.Get("/{customerId}", customerHandler.CustomerGet)
		r.Post("/", customerHandler.CustomerPost)
	})
}
