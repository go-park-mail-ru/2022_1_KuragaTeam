package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/internal"
	"myapp/internal/movie"
)

type movieStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) movie.Storage {
	return &movieStorage{db: db}
}

func (ms *movieStorage) GetOne(id int) (*internal.Movie, error) {
	sql := "SELECT id, name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
		"picture, video, trailer FROM movies WHERE id=$1"

	var selectedMovie internal.Movie
	err := ms.db.QueryRow(context.Background(), sql, id).Scan(&selectedMovie.ID, &selectedMovie.Name,
		&selectedMovie.NamePicture, &selectedMovie.Year, &selectedMovie.Duration, &selectedMovie.AgeLimit,
		&selectedMovie.Description, &selectedMovie.KinopoiskRating, &selectedMovie.Tagline, &selectedMovie.Picture,
		&selectedMovie.Video, &selectedMovie.Trailer)
	if err != nil {
		return nil, err
	}

	return &selectedMovie, nil
}

func (ms *movieStorage) GetAllMovies(limit, offset int) ([]internal.Movie, error) {
	sql := "SELECT id, name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
		"picture, video, trailer FROM movies LIMIT $1 OFFSET $2"

	selectedMovies := make([]internal.Movie, 0, limit)

	rows, err := ms.db.Query(context.Background(), sql, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var singleMovie internal.Movie
		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.NamePicture, &singleMovie.Year,
			&singleMovie.Duration, &singleMovie.AgeLimit, &singleMovie.Description, &singleMovie.KinopoiskRating,
			&singleMovie.Tagline, &singleMovie.Picture, &singleMovie.Video, &singleMovie.Trailer); err != nil {
			return nil, err
		}
		selectedMovies = append(selectedMovies, singleMovie)
	}

	return selectedMovies, nil
}

func (ms *movieStorage) GetRandomMovie() (*internal.MainMovieInfoDTO, error) {
	sql := "SELECT id, name, tagline, picture FROM movies ORDER BY RANDOM() LIMIT 1"

	var mainMovie internal.MainMovieInfoDTO
	err := ms.db.QueryRow(context.Background(), sql).Scan(&(mainMovie.ID), &(mainMovie.Name),
		&(mainMovie.Tagline), &(mainMovie.Picture))
	if err != nil {
		return nil, err
	}

	return &mainMovie, nil
}
