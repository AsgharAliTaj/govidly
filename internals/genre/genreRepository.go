package genre

import "github.com/jmoiron/sqlx"

type GenreRepository interface {
	GetGenre() (Genre, error)
	CreateGenre(Genre) error
	DeleteGenre() error
	GetAllGenre() ([]Genre, error)
}

type genreRepository struct {
	databse *sqlx.DB
}

func NewGenreRepository(databse *sqlx.DB) GenreRepository {
	return &genreRepository{databse: databse}
}

func (g *genreRepository) GetGenre() (genre Genre, err error) {
	return genre, nil
}

func (g *genreRepository) CreateGenre(genre Genre) (err error) {
	err = g.databse.Get(
		&genre,
		"INSERT INTO genres (id, name) VALUES ($1, $2) RETURNING *",
		genre.ID,
		genre.Name,
	)
	if err != nil {
		return err
	}
	return nil
}

func (g *genreRepository) DeleteGenre() (err error) {
	return nil
}

func (g *genreRepository) GetAllGenre() (genres []Genre, err error) {
	return nil, nil
}
