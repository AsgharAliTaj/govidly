package auth

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/asgharalitaj/govidly/internals/user"
)

type Auther interface {
	GetUser(string) (user.User, error)
}

type authRepository struct {
	database *sqlx.DB
}

func (a *authRepository) GetUser(email string) (user user.User, err error) {
	err = a.database.Get(&user, "SELECT * from users where email=$1", email)
	if err != nil {
		return user, errors.New("invalid email or password")
	}
	return user, nil
}
