package repository

import (
	"database/sql"
	"myapp/internal/genre"
)

type genreStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) genre.Storage {
	return &genreStorage{db: db}
}

func (ms *genreStorage) GetByMovieID(id int) ([]string, error) {
	genres := make([]string, 0)

	sqlScript := "SELECT g.name FROM genre AS g JOIN movies_genre mv_g ON mv_g.genre_id = g.id " +
		"WHERE mv_g.movie_id = $1 ORDER BY mv_g.id"
	rows, err := ms.db.Query(sqlScript, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nextGenre string
		if err = rows.Scan(&nextGenre); err != nil {
			return nil, err
		}
		genres = append(genres, nextGenre)
	}

	return genres, nil
}

//func (ms *genreStorage) GetByID(id int) (string, error) {
//	var selectedGenre string
//
//	sqlScript := "SELECT name FROM genre WHERE id = $1"
//	err := ms.db.QueryRow(sqlScript, id).Scan(&selectedGenre)
//	if err != nil {
//		return "", err
//	}
//
//	return selectedGenre, nil
//}
