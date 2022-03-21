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
	sql := "SELECT DISTINCT m.id, m.name, m.picture FROM movies AS m " +
		"JOIN movies_genre m_g ON m_g.movie_id = m.id " +
		"JOIN movies_genre m_g2 ON m_g2.genre_id = m_g.genre_id " +
		"WHERE m_g2.movie_id=$1"

	var selectedMC moviesCompilations.MovieCompilation
	selectedMC.Name = "Похожие по жанру"

	rows, err := ms.db.Query(context.Background(), sql, movieID)

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
	sql := "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_staff m_s ON " +
		"m_s.movie_id = m.id WHERE m_s.person_id=$1"

	var selectedMovieCompilation moviesCompilations.MovieCompilation
	selectedMovieCompilation.Name = "Фильмография"
	rows, err := ms.db.Query(context.Background(), sql, personID)

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
	return moviesCompilations.MovieCompilation{}, nil
}
func (ms *movieCompilationsStorage) GetTopByYear(year int) (moviesCompilations.MovieCompilation, error) {
	return moviesCompilations.MovieCompilation{}, nil
}
