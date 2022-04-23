package repository

import (
	"database/sql"
	"myapp/internal/microservices/movie"
	"myapp/internal/microservices/movie/proto"
)

type movieStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) movie.Storage {
	return &movieStorage{db: db}
}

func (ms *movieStorage) GetOne(id int) (*proto.Movie, error) {
	sqlScript := "SELECT id, name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
		"picture, video, trailer FROM movies WHERE id=$1"

	var selectedMovie proto.Movie
	err := ms.db.QueryRow(sqlScript, id).Scan(&selectedMovie.ID, &selectedMovie.Name,
		&selectedMovie.NamePicture, &selectedMovie.Year, &selectedMovie.Duration, &selectedMovie.AgeLimit,
		&selectedMovie.Description, &selectedMovie.KinopoiskRating, &selectedMovie.Tagline, &selectedMovie.Picture,
		&selectedMovie.Video, &selectedMovie.Trailer)
	if err != nil {
		return nil, err
	}

	return &selectedMovie, nil
}

func (ms *movieStorage) GetAllMovies(limit, offset int) ([]*proto.Movie, error) {
	sqlScript := "SELECT id, name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
		"picture, video, trailer FROM movies LIMIT $1 OFFSET $2"

	selectedMovies := make([]*proto.Movie, 0, limit)

	rows, err := ms.db.Query(sqlScript, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var singleMovie proto.Movie
		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.NamePicture, &singleMovie.Year,
			&singleMovie.Duration, &singleMovie.AgeLimit, &singleMovie.Description, &singleMovie.KinopoiskRating,
			&singleMovie.Tagline, &singleMovie.Picture, &singleMovie.Video, &singleMovie.Trailer); err != nil {
			return nil, err
		}
		selectedMovies = append(selectedMovies, &singleMovie)
	}

	return selectedMovies, nil
}

func (ms *movieStorage) GetRandomMovie() (*proto.MainMovie, error) {
	sqlScript := "SELECT id, name_picture, tagline, picture FROM movies ORDER BY RANDOM() LIMIT 1"

	var mainMovie proto.MainMovie
	err := ms.db.QueryRow(sqlScript).Scan(&(mainMovie.ID), &(mainMovie.NamePicture),
		&(mainMovie.Tagline), &(mainMovie.Picture))
	if err != nil {
		return nil, err
	}

	return &mainMovie, nil
}
