package movie

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Movier interface {
	GetAllMovies() ([]Movie, error)
	GetMovie(string) (Movie, error)
	CreateMovie(Movie) error
	DeleteMovie(string) error
}

type movieRepository struct {
	databse *sqlx.DB
}

func (m *movieRepository) GetAllMovies() (movies []Movie, err error) {
	genresName := make(map[uuid.UUID][]Moviegenre)

	err = m.databse.Select(&movies, "select * from movies")
	if err != nil {
		return nil, err
	}

	rows, err := m.databse.Queryx(
		"select movie_id, genre_name  from movies_genres",
	)
	for rows.Next() {
		var (
			movieId   uuid.UUID
			genreName string
		)
		rows.Scan(&movieId, &genreName)
		if value, ok := genresName[movieId]; ok {
			genresName[movieId] = append(value, Moviegenre{Name: genreName})
		} else {
			genresName[movieId] = []Moviegenre{{Name: genreName}}
		}

	}
	if err != nil {
		return nil, err
	}
	for key, value := range genresName {
		for i := 0; i < len(movies); i++ {
			if key == movies[i].ID {
				movies[i].Genres = value
			}
		}
	}
	return
}

func (m *movieRepository) GetMovie(id string) (movie Movie, err error) {
	err = m.databse.Get(&movie, "SELECT * FROM movies where id  = $1 ", id)
	if err != nil {
		return movie, err
	}
	return
}

func (m *movieRepository) CreateMovie(movie Movie) (err error) {
	return nil
}

func (m *movieRepository) DeleteMovie(id string) (err error) {
	return nil
}

// genresName := make(map[uuid.UUID][]Moviegenre)
//
// 	rows, err := m.databse.Queryx(
// 		"select movies.id, movies.name, movies.number_in_stock, movies.daily_rental_rate, movies_genres.genre_name from movies inner join movies_genres on movies.id = movies_genres.movie_id",
// 	)
// 	for rows.Next() {
// 		var (
// 			movieId         uuid.UUID
// 			movieName       string
// 			numberInStock   int
// 			dailyRentalRate int
// 			genreName       string
// 		)
//
// 		rows.Scan(&movieId, &movieName, &numberInStock, &dailyRentalRate, &genreName)
// 		if value, ok := genresName[movieId]; ok {
// 			genresName[movieId] = append(value, Moviegenre{Name: genreName})
// 		} else {
// 			genresName[movieId] = []Moviegenre{{Name: genreName}}
// 		}
//
// 		for key, value := range genresName {
// 			movies = append(movies, Movie{
// 				ID:              key,
// 				Name:            movieName,
// 				NumberInStock:   numberInStock,
// 				DailyRentalRate: dailyRentalRate,
// 				Genres:          value,
// 			})
// 		}
// 	}
