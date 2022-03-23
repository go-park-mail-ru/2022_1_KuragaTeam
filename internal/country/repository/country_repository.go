package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/internal/country"
)

type countryStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) country.Storage {
	return &countryStorage{db: db}
}

func (ms *countryStorage) GetByMovieID(id int) ([]string, error) {
	countries := make([]string, 0)

	sql := "SELECT c.name FROM country AS c JOIN movies_countries mv_c ON mv_c.country_id = c.id " +
		"WHERE mv_c.movie_id = $1"
	rows, err := ms.db.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nextCountry string
		if err = rows.Scan(&nextCountry); err != nil {
			return nil, err
		}
		countries = append(countries, nextCountry)
	}

	return countries, nil
}
