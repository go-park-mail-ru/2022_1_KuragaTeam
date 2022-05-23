package repository

import (
	"database/sql"
)

type positionStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *positionStorage {
	return &positionStorage{db: db}
}

func (ms *positionStorage) GetByPersonID(id int) ([]string, error) {
	personPositions := make([]string, 0)

	sqlScript := "SELECT pos.name FROM position AS pos JOIN movies_staff mv_s ON mv_s.position_id = pos.id " +
		"WHERE mv_s.person_id = $1"
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
		var personPosition string
		if err = rows.Scan(&personPosition); err != nil {
			return nil, err
		}
		personPositions = append(personPositions, personPosition)
	}

	return personPositions, nil
}

func (ms *positionStorage) GetPersonPosByMovieID(personID, movieID int) (string, error) {
	var positionName string
	sqlScript := "SELECT pos.name FROM position AS pos JOIN movies_staff mv_s ON mv_s.position_id = pos.id " +
		"WHERE mv_s.movie_id = $1 AND WHERE mv_s.person_id = $2"
	err := ms.db.QueryRow(sqlScript, movieID, personID).Scan(&positionName)
	if err != nil {
		return "", err
	}
	return positionName, nil
}
