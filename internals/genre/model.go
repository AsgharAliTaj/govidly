package genre

import (
	"errors"

	"github.com/google/uuid"
)

// import "github.com/google/uuid"

type Genre struct {
	ID   uuid.UUID `json:"id"   db:"id"`
	Name string    `json:"name" db:"name"`
}

func validateName(g Genre) error {
	if g.Name == "" && len(g.Name) > 64 {
		return errors.New("name is either empty or exceed the maximum length")
	}
	return nil
}
