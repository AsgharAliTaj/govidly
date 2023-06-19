package genre

import (
	"errors"
)

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name" db:"name"`
}

func validateName(g Genre) error {
	if g.Name == "" && len(g.Name) > 64 {
		return errors.New("name is either empty or exceed the maximum length")
	}
	return nil
}
