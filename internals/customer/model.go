package customer

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID     uuid.UUID `db:"id" json:"id"`
	Name   string    `db:"name" json:"name"`
	phone  string    `db:"phone" json: "phone"`
	IsGold bool      `db:"is_gold" json: "isGold"`
}
