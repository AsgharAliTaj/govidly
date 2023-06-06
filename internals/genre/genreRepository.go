package genre

import (
	"github.com/jmoiron/sqlx"
)

type Genrer interface {
	GetAllGenre() ([]Genre, error)
	GetGenre(int) (Genre, error)
	CreateGenre(Genre) error
	DeleteGenre(int) error
}

type genreRepository struct {
	database *sqlx.DB
}

// get all genres
func (g *genreRepository) GetAllGenre() (genres []Genre, err error) {
	err = g.database.Select(&genres, "SELECT * FROM genres ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	return genres, nil
}

// get a single genre
func (g *genreRepository) GetGenre(id int) (genre Genre, err error) {
	err = g.database.Get(&genre, "SELECT * from genres where id=$1", id)
	if err != nil {
		return genre, err
	}
	return genre, nil
}

// create a genre
func (g *genreRepository) CreateGenre(genre Genre) (err error) {
	err = g.database.Get(
		&genre,
		"INSERT INTO genres (name) VALUES ($1) RETURNING *",
		genre.Name,
	)
	if err != nil {
		return err
	}
	return nil
}

// delete a genre
func (g *genreRepository) DeleteGenre(id int) (err error) {
	_, err = g.database.Exec("DELETE from genres where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
