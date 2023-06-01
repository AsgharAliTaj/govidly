package genre

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type GenreRepository interface {
	GetGenre(int) (Genre, error)
	CreateGenre(Genre) error
	DeleteGenre(int) error
	GetAllGenre() ([]Genre, error)
}

type genreRepository struct {
	databse *sqlx.DB
}

func NewGenreRepository(databse *sqlx.DB) GenreRepository {
	return &genreRepository{databse: databse}
}

func (g *genreRepository) GetGenre(id int) (genre Genre, err error) {
	err = g.databse.Get(&genre, "SELECT * from genres where id=$1", id)
	if err != nil {
		return genre, err
	}
	return genre, nil
}

func (g *genreRepository) CreateGenre(genre Genre) (err error) {
	err = g.databse.Get(
		&genre,
		"INSERT INTO genres (name) VALUES ($1) RETURNING *",
		genre.Name,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (g *genreRepository) DeleteGenre(id int) (err error) {
	_, err = g.databse.Exec("DELETE from genres where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (g *genreRepository) GetAllGenre() (genres []Genre, err error) {
	err = g.databse.Select(&genres, "SELECT * FROM genres")
	if err != nil {
		return nil, err
	}
	return genres, nil
}
