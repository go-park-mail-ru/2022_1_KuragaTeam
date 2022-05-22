package repository

import (
	"database/sql"
	"myapp/internal"
)

type genreStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *genreStorage {
	return &genreStorage{db: db}
}

func (ms *genreStorage) GetByMovieID(id int) ([]internal.Genre, error) {
	genres := make([]internal.Genre, 0)

	sqlScript := "SELECT g.id, g.name FROM genre AS g JOIN movies_genre mv_g ON mv_g.genre_id = g.id " +
		"WHERE mv_g.movie_id = $1 ORDER BY mv_g.id"
	rows, err := ms.db.Query(sqlScript, id)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nextGenre internal.Genre
		if err = rows.Scan(&nextGenre.ID, &nextGenre.Name); err != nil {
			return nil, err
		}
		genres = append(genres, nextGenre)
	}

	return genres, nil
}
