package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	"myapp/internal/moviesCompilations"
)

type movieCompilationsStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) moviesCompilations.Storage {
	return &movieCompilationsStorage{db: db}
}

func (ms *movieCompilationsStorage) GetByGenre(genreID int) (moviesCompilations.MovieCompilation, error) {
	sql := "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
		"m_g.movie_id = m.id WHERE m_g.genre_id=$1"

	sqlGenre := "SELECT name FROM genre WHERE id=$1"

	var selectedMovieCompilation moviesCompilations.MovieCompilation

	err := ms.db.QueryRow(context.Background(), sqlGenre, genreID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}

	rows, err := ms.db.Query(context.Background(), sql, genreID)

	for rows.Next() {
		var selectedMovie moviesCompilations.Movie
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return moviesCompilations.MovieCompilation{}, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, selectedMovie)
	}
	return selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByMovie(movieID int) (moviesCompilations.MovieCompilation, error) {
	return moviesCompilations.MovieCompilation{}, nil
}
func (ms *movieCompilationsStorage) GetByPerson(personID int) (moviesCompilations.MovieCompilation, error) {
	return moviesCompilations.MovieCompilation{}, nil
}
func (ms *movieCompilationsStorage) GetTop(limit int) (moviesCompilations.MovieCompilation, error) {
	return moviesCompilations.MovieCompilation{}, nil
}
func (ms *movieCompilationsStorage) GetTopByYear(year int) (moviesCompilations.MovieCompilation, error) {
	return moviesCompilations.MovieCompilation{}, nil
}

//func (ms *movieCompilationsStorage) GetOne(id int) (*moviesCompilations.Movie, error) {
//	sql := "SELECT id, name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
//		"picture, video, trailer FROM movies WHERE id=$1"
//
//	var selectedMovie internal.Movie
//	err := ms.db.QueryRow(context.Background(), sql, id).Scan(&selectedMovie.ID, &selectedMovie.Name,
//		&selectedMovie.NamePicture, &selectedMovie.Year, &selectedMovie.Duration, &selectedMovie.AgeLimit,
//		&selectedMovie.Description, &selectedMovie.KinopoiskRating, &selectedMovie.Tagline, &selectedMovie.Picture,
//		&selectedMovie.Video, &selectedMovie.Trailer)
//	if err != nil {
//		return nil, err
//	}
//
//	return &selectedMovie, nil
//}
//
//func (ms *movieCompilationsStorage) GetAllMovies(limit, offset int) ([]internal.Movie, error) {
//	sql := "SELECT id, name, name_picture, year, duration, age_limit, description, kinopoisk_rating, tagline, " +
//		"picture, video, trailer FROM movies LIMIT $1 OFFSET $2"
//
//	selectedMovies := make([]internal.Movie, 0, limit)
//
//	rows, err := ms.db.Query(context.Background(), sql, limit, offset)
//	if err != nil {
//		return nil, err
//	}
//
//	for rows.Next() {
//		var singleMovie internal.Movie
//		if err = rows.Scan(&singleMovie.ID, &singleMovie.Name, &singleMovie.NamePicture, &singleMovie.Year,
//			&singleMovie.Duration, &singleMovie.AgeLimit, &singleMovie.Description, &singleMovie.KinopoiskRating,
//			&singleMovie.Tagline, &singleMovie.Picture, &singleMovie.Video, &singleMovie.Trailer); err != nil {
//			return nil, err
//		}
//		selectedMovies = append(selectedMovies, singleMovie)
//	}
//
//	return selectedMovies, nil
//}
//
//func (ms *movieCompilationsStorage) GetRandomMovie() (*internal.MainMovieInfoDTO, error) {
//	sql := "SELECT id, name, tagline, picture FROM movies ORDER BY RANDOM() LIMIT 1"
//
//	var mainMovie internal.MainMovieInfoDTO
//	err := ms.db.QueryRow(context.Background(), sql).Scan(&(mainMovie.ID), &(mainMovie.Name),
//		&(mainMovie.Tagline), &(mainMovie.Picture))
//	if err != nil {
//		return nil, err
//	}
//
//	return &mainMovie, nil
//}
