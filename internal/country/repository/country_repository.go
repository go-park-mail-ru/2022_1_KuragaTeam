package repository

import (
	"database/sql"
	"myapp/internal/country"
)

type countryStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) country.Storage {
	return &countryStorage{db: db}
}

func (ms *countryStorage) GetByMovieID(id int) ([]string, error) {
	countries := make([]string, 0)

	sqlScript := "SELECT c.name FROM country AS c JOIN movies_countries mv_c ON mv_c.country_id = c.id " +
		"WHERE mv_c.movie_id = $1 ORDER BY mv_c.id"
	rows, err := ms.db.Query(sqlScript, id)
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
