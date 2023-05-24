package genre

import (
	"github.com/google/uuid"
)

type Genre struct {
	ID   uuid.UUID `json:"id",db:"id"`
	Name string    `db:name`
}

// Valide Genre functions
