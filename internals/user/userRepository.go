package user

import "github.com/jmoiron/sqlx"

type Userer interface {
	CreateUser(User) error
	GetUser(string) (User, error)
}

type userRepository struct {
	database *sqlx.DB
}

func (u *userRepository) GetUser(email string) (user User, err error) {
	err = u.database.Get(&user, "SELECT * from users where email=$1", email)
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
