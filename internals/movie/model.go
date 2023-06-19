package movie

import (
	"github.com/google/uuid"
)

type Moviegenre struct {
	Name string `json:"name"`
}

type Movie struct {
	ID              uuid.UUID    `json:"id"              db:"id"`
	Name            string       `json:"name"            db:"name"`
	NumberInStock   int          `json:"numberInStock"   db:"number_in_stock"`
	DailyRentalRate int          `json:"dailyRentalRate" db:"daily_rental_rate"`
	Genres          []Moviegenre `json:"genres"`
}
