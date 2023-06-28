package user

import (
	"errors"
	"net/mail"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func validateEmail(address string) error {
	if len(address) > 245 {
		return errors.New("Email length exceeds 245 characters")
	}
	_, err := mail.ParseAddress(address)
	if err != nil {
		return errors.New("Email is invalid")
	}
	return nil
}
