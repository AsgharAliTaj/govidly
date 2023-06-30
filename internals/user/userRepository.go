package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Userer interface {
	CreateUser(User) error
	GetUser(uuid.UUID) (User, error)
}

type userRepository struct {
	database *sqlx.DB
}

func (u *userRepository) GetUser(id uuid.UUID) (user User, err error) {
	err = u.database.Get(&user, "SELECT id, name, email from users where id=$1", id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) CreateUser(user User) (err error) {
	err = u.database.Get(
		&user,
		"INSERT INTO users (id,name,email,Password) VALUES ($1, $2, $3, $4) RETURNING *",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}
