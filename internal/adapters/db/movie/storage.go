package movie

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/internal/domain"

	"myapp/internal/domain/movie"
)

type movieStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) movie.Storage {
	return &movieStorage{db: db}
}

func (ms *movieStorage) GetOne(id int) (*domain.Movie, error) {
	sql := "SELECT name, description, picture, video, trailer FROM movies WHERE id=$1"

	var selectedMovie domain.Movie

	err := ms.db.QueryRow(context.Background(), sql, id).Scan(&selectedMovie)
	if err != nil {
		return nil, err
	}

	return &selectedMovie, nil
}

func (ms *movieStorage) GetRandom(limit int) ([]domain.Movie, error) {
	sql := "SELECT id, name, description, picture, video, trailer FROM movies " +
		"ORDER BY RANDOM() LIMIT $1"

	selectedMovies := make([]domain.Movie, 0, limit)

	rows, err := ms.db.Query(context.Background(), sql, limit)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var singleMovie domain.Movie
		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.Description, &singleMovie.Picture,
			&singleMovie.Video, &singleMovie.Trailer); err != nil {
			return nil, err
		}
		sql2 := "SELECT g.name FROM genre AS g JOIN movies_genre mv_g ON mv_g.genre_id = g.id " +
			"WHERE mv_g.movie_id = $1"
		rows2, err2 := ms.db.Query(context.Background(), sql2, singleMovie.ID)
		if err2 != nil {
			singleMovie.Genre = append(singleMovie.Genre, err.Error())
		}
		for rows2.Next() {
			var genre string
			if err2 = rows2.Scan(&genre); err2 != nil {
				singleMovie.Genre = append(singleMovie.Genre, err.Error())
				break
			}
			singleMovie.Genre = append(singleMovie.Genre, genre)
		}

		selectedMovies = append(selectedMovies, singleMovie)
	}

	return selectedMovies, nil
}
