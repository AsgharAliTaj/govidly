package user

import (
	"errors"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MyCustomClaim struct {
	Id uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u *User) GenerateAuthToken() (string, error) {
	claims := MyCustomClaim{
		u.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		return "", err
	}
	return ss, nil
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
