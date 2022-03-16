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
	sql := "SELECT id, name, year, description, picture, video, trailer FROM movies WHERE id=$1"

	var selectedMovie domain.Movie
	err := ms.db.QueryRow(context.Background(), sql, id).Scan(&selectedMovie.ID, &selectedMovie.Name,
		&selectedMovie.Year, &selectedMovie.Description, &selectedMovie.Picture,
		&selectedMovie.Video, &selectedMovie.Trailer)
	if err != nil {
		return nil, err
	}

	return &selectedMovie, nil
}

func (ms *movieStorage) GetRandomMovies(limit, offset int) ([]domain.Movie, error) {
	sql := "SELECT id, name, year, description, picture, video, trailer FROM movies LIMIT $1 OFFSET $2"

	selectedMovies := make([]domain.Movie, 0, limit)

	rows, err := ms.db.Query(context.Background(), sql, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var singleMovie domain.Movie
		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.Year, &singleMovie.Description,
			&singleMovie.Picture, &singleMovie.Video, &singleMovie.Trailer); err != nil {
			return nil, err
		}
		selectedMovies = append(selectedMovies, singleMovie)
	}

	return selectedMovies, nil
}

func (ms *movieStorage) GetRandomMovie() (*domain.MainMovieInfoDTO, error) {
	sql := "SELECT id, name, picture FROM movies ORDER BY RANDOM() LIMIT 1"

	var mainMovie domain.MainMovieInfoDTO
	err := ms.db.QueryRow(context.Background(), sql).Scan(&(mainMovie.ID), &(mainMovie.Name), &(mainMovie.Picture))
	if err != nil {
		return nil, err
	}

	return &mainMovie, nil
}
