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
	sql := "SELECT id, name, description, picture, video, trailer FROM movies ORDER BY RANDOM() LIMIT $1"

	selectedMovies := make([]domain.Movie, 0, limit)

	rows, err := ms.db.Query(context.Background(), sql, limit)
	if err != nil {
		return nil, err
	}
	var singleMovie domain.Movie
	for rows.Next() {
		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.Description, &singleMovie.Picture,
			&singleMovie.Video, &singleMovie.Trailer); err != nil {
			return nil, err
		}
		selectedMovies = append(selectedMovies, singleMovie)
	}

	return selectedMovies, nil
}
