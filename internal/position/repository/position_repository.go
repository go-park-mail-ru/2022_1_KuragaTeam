package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/internal/position"
)

type positionStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) position.Storage {
	return &positionStorage{db: db}
}

func (ms *positionStorage) GetByPersonID(id int) ([]string, error) {
	personPositions := make([]string, 0)

	sql := "SELECT pos.name FROM position AS pos JOIN movies_staff mv_s ON mv_s.position_id = pos.id " +
		"WHERE mv_s.person_id = $1"
	rows, err := ms.db.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var personPosition string
		if err = rows.Scan(&personPosition); err != nil {
			return nil, err
		}
		personPositions = append(personPositions, personPosition)
	}

	return personPositions, nil
}

func (ms *positionStorage) GetPersonPosByMovieID(personID, movieId int) (string, error) {
	var positionName string
	sql := "SELECT pos.name FROM position AS pos JOIN movies_staff mv_s ON mv_s.position_id = pos.id " +
		"WHERE mv_s.movie_id = $1 AND WHERE mv_s.person_id = $2"
	err := ms.db.QueryRow(context.Background(), sql, movieId, personID).Scan(&positionName)
	if err != nil {
		return "", err
	}
	return positionName, nil
}
