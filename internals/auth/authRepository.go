package auth

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type Auther interface {
	GetUser(string) (UserAuth, error)
}

type authRepository struct {
	database *sqlx.DB
}

func (a *authRepository) GetUser(email string) (user UserAuth, err error) {
	err = a.database.Get(&user, "SELECT * from users where email=$1", email)
	if err != nil {
		return user, errors.New("invalid email or password")
	}
	return user, nil
}
