package repository

import (
	"database/sql"
	"fmt"
	"myapp/internal/moviesCompilations"
)

type movieCompilationsStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) moviesCompilations.Storage {
	return &movieCompilationsStorage{db: db}
}

const (
	getByGenreSQL = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
		"m_g.movie_id = m.id WHERE m_g.genre_id=$1"
	getGenreNameSQL = "SELECT name FROM genre WHERE id=$1"
	getByCountrySQL = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_countries m_c ON " +
		"m_c.movie_id = m.id WHERE m_c.country_id=$1"
	getCountryNameSQL = "SELECT name FROM country WHERE id=$1"
	getByMovieSQL     = "SELECT DISTINCT m.id, m.name, m.picture FROM movies AS m " +
		"JOIN movies_genre m_g ON m_g.movie_id = m.id " +
		"JOIN movies_genre m_g2 ON m_g2.genre_id = m_g.genre_id " +
		"WHERE m_g2.movie_id=$1"
	getByPersonSQL = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_staff m_s ON " +
		"m_s.movie_id = m.id WHERE m_s.person_id=$1"
	getTopSQL       = "SELECT id, name, picture FROM movies ORDER BY kinopoisk_rating DESC LIMIT $1"
	getTopByYearSQL = "SELECT id, name, picture FROM movies WHERE year=$1 ORDER BY kinopoisk_rating DESC"
)

func (ms *movieCompilationsStorage) GetByGenre(genreID int) (moviesCompilations.MovieCompilation, error) {
	var selectedMovieCompilation moviesCompilations.MovieCompilation

	err := ms.db.QueryRow(getGenreNameSQL, genreID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}

	rows, err := ms.db.Query(getByGenreSQL, genreID)
	defer rows.Close()

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

func (ms *movieCompilationsStorage) GetByCountry(countryID int) (moviesCompilations.MovieCompilation, error) {
	var selectedMovieCompilation moviesCompilations.MovieCompilation

	err := ms.db.QueryRow(getCountryNameSQL, countryID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return moviesCompilations.MovieCompilation{}, err
	}

	rows, err := ms.db.Query(getByCountrySQL, countryID)
	defer rows.Close()

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
	var selectedMC moviesCompilations.MovieCompilation
	selectedMC.Name = "Похожие по жанру"

	rows, err := ms.db.Query(getByMovieSQL, movieID)
	defer rows.Close()

	for rows.Next() {
		var selectedMovie moviesCompilations.Movie
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return moviesCompilations.MovieCompilation{}, err
		}
		// Костыль запроса. Необходимо добавить в запрос исключение исходного фильма
		if selectedMovie.ID != movieID {
			selectedMC.Movies = append(selectedMC.Movies, selectedMovie)
		}
	}
	return selectedMC, nil
}

func (ms *movieCompilationsStorage) GetByPerson(personID int) (moviesCompilations.MovieCompilation, error) {

	var selectedMovieCompilation moviesCompilations.MovieCompilation
	selectedMovieCompilation.Name = "Фильмография"
	rows, err := ms.db.Query(getByPersonSQL, personID)
	defer rows.Close()

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

func (ms *movieCompilationsStorage) GetTop(limit int) (moviesCompilations.MovieCompilation, error) {

	var selectedMovieCompilation moviesCompilations.MovieCompilation
	selectedMovieCompilation.Name = "Топ рейтинга"
	rows, err := ms.db.Query(getTopSQL, limit)
	defer rows.Close()

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

func (ms *movieCompilationsStorage) GetTopByYear(year int) (moviesCompilations.MovieCompilation, error) {

	var selectedMovieCompilation moviesCompilations.MovieCompilation
	selectedMovieCompilation.Name = fmt.Sprintf("Лучшее за %d год", year)
	rows, err := ms.db.Query(getTopByYearSQL, year)
	defer rows.Close()

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
