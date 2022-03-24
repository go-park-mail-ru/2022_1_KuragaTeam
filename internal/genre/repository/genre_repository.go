package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/internal/genre"
)

type genreStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) genre.Storage {
	return &genreStorage{db: db}
}

func (ms *genreStorage) GetByMovieID(id int) ([]string, error) {
	genres := make([]string, 0)

	sql := "SELECT g.name FROM genre AS g JOIN movies_genre mv_g ON mv_g.genre_id = g.id " +
		"WHERE mv_g.movie_id = $1 ORDER BY mv_g.id"
	rows, err := ms.db.Query(context.Background(), sql, id)
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
