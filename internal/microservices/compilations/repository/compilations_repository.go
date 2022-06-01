package repository

import (
	"database/sql"
	"fmt"
	"myapp/internal/constants"
	"myapp/internal/microservices/compilations/proto"
)

type movieCompilationsStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *movieCompilationsStorage {
	return &movieCompilationsStorage{db: db}
}

const (
	getAllMoviesSQL = "SELECT m.id, m.name, m.picture FROM movies AS m WHERE m.is_movie=$1 LIMIT $2 OFFSET $3"
	getByGenreSQL   = "SELECT m.id, m.name, m.picture FROM movies AS m JOIN movies_genre m_g ON " +
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
	getTopSQL          = "SELECT id, name, picture FROM movies ORDER BY movies.rating_sum/movies.rating_count DESC LIMIT $1"
	getTopByYearSQL    = "SELECT id, name, picture FROM movies WHERE year=$1 ORDER BY kinopoisk_rating DESC"
	getFavorites       = "SELECT id, name, picture, is_movie FROM movies WHERE id=$1;"
	findMovie          = "SELECT id, name, picture FROM movies where to_tsvector('russian', name) @@ to_tsquery('russian', $1) AND is_movie=$2 ORDER BY name LIMIT $3;"
	findMovieByPartial = "SELECT id, name, picture FROM movies where name ILIKE $1 AND is_movie=$2 ORDER BY name LIMIT $3;"
)

func (ms *movieCompilationsStorage) GetAllMovies(limit, offset int, isMovie bool) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	rows, err := ms.db.Query(getAllMoviesSQL, isMovie, limit, offset)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByGenre(genreID int) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	err := ms.db.QueryRow(getGenreNameSQL, genreID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return nil, err
	}

	rows, err := ms.db.Query(getByGenreSQL, genreID)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByCountry(countryID int) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	err := ms.db.QueryRow(getCountryNameSQL, countryID).Scan(&selectedMovieCompilation.Name)
	if err != nil {
		return nil, err
	}

	rows, err := ms.db.Query(getByCountrySQL, countryID)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetByMovie(movieID int) (*proto.MovieCompilation, error) {
	var selectedMC proto.MovieCompilation
	selectedMC.Name = "Похожие по жанру"

	rows, err := ms.db.Query(getByMovieSQL, movieID)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		// Костыль запроса. Необходимо добавить в запрос исключение исходного фильма
		if selectedMovie.ID != int64(movieID) {
			selectedMC.Movies = append(selectedMC.Movies, &selectedMovie)
		}
	}
	return &selectedMC, nil
}

func (ms *movieCompilationsStorage) getCompilation(sqlQuery string, args int) (*proto.MovieCompilation, error) {

	var selectedMovieCompilation proto.MovieCompilation
	rows, err := ms.db.Query(sqlQuery, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil

}

func (ms *movieCompilationsStorage) GetByPerson(personID int) (*proto.MovieCompilation, error) {
	selectedMovieCompilation, err := ms.getCompilation(getByPersonSQL, personID)
	if err != nil {
		return nil, err
	}
	selectedMovieCompilation.Name = "Фильмография"
	return selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetTop(limit int) (*proto.MovieCompilation, error) {
	selectedMovieCompilation, err := ms.getCompilation(getTopSQL, limit)
	if err != nil {
		return nil, err
	}
	selectedMovieCompilation.Name = "Топ рейтинга"
	return selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetTopByYear(year int) (*proto.MovieCompilation, error) {
	selectedMovieCompilation, err := ms.getCompilation(getTopByYearSQL, year)
	if err != nil {
		return nil, err
	}
	selectedMovieCompilation.Name = fmt.Sprintf("Лучшее за %d год", year)
	return selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetFavorites(data *proto.GetFavoritesOptions) (*proto.MovieCompilationsArr, error) {
	selectedMovieCompilation := &proto.MovieCompilationsArr{}
	compilation := &proto.MovieCompilation{Name: "Фильмы"}
	selectedMovieCompilation.MovieCompilations = append(selectedMovieCompilation.MovieCompilations, compilation)
	compilation = &proto.MovieCompilation{Name: "Сериалы"}
	selectedMovieCompilation.MovieCompilations = append(selectedMovieCompilation.MovieCompilations, compilation)

	for _, id := range data.Id {
		rows, err := ms.db.Query(getFavorites, id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var selectedMovie proto.MovieInfo
			var isMovie bool
			err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture, &isMovie)
			if err != nil {
				return nil, err
			}

			if isMovie {
				selectedMovieCompilation.MovieCompilations[0].Movies = append(selectedMovieCompilation.MovieCompilations[0].Movies, &selectedMovie)
			} else {
				selectedMovieCompilation.MovieCompilations[1].Movies = append(selectedMovieCompilation.MovieCompilations[1].Movies, &selectedMovie)
			}
		}
	}

	return selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) FindMovie(text string, isMovie bool) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	rows, err := ms.db.Query(findMovie, text, isMovie, constants.MoviesSearchLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) FindMovieByPartial(text string, isMovie bool) (*proto.MovieCompilation, error) {
	var selectedMovieCompilation proto.MovieCompilation

	rows, err := ms.db.Query(findMovieByPartial, "%"+text+"%", isMovie, constants.MoviesSearchLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedMovie proto.MovieInfo
		err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture)
		if err != nil {
			return nil, err
		}
		selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
	}
	return &selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetFavoritesFilms(data *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	selectedMovieCompilation := &proto.MovieCompilation{Name: "Фильмы"}

	for _, id := range data.Id {
		rows, err := ms.db.Query(getFavorites, id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var selectedMovie proto.MovieInfo
			var isMovie bool
			err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture, &isMovie)
			if err != nil {
				return nil, err
			}

			if isMovie {
				selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
			}
		}
	}

	return selectedMovieCompilation, nil
}

func (ms *movieCompilationsStorage) GetFavoritesSeries(data *proto.GetFavoritesOptions) (*proto.MovieCompilation, error) {
	selectedMovieCompilation := &proto.MovieCompilation{Name: "Сериалы"}

	for _, id := range data.Id {
		rows, err := ms.db.Query(getFavorites, id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var selectedMovie proto.MovieInfo
			var isMovie bool
			err = rows.Scan(&selectedMovie.ID, &selectedMovie.Name, &selectedMovie.Picture, &isMovie)
			if err != nil {
				return nil, err
			}

			if isMovie {
				selectedMovieCompilation.Movies = append(selectedMovieCompilation.Movies, &selectedMovie)
			}
		}
	}

	return selectedMovieCompilation, nil
}
