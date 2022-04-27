package repository

import (
	"database/sql"
	"myapp/internal"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/persons"
)

type staffStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) persons.Storage {
	return &staffStorage{db: db}
}

const (
	sqlGetByMovieID = "SELECT p.id, p.name, p.photo, pos.name FROM person AS p JOIN movies_staff mv_s ON mv_s.person_id = p.id " +
		"JOIN position pos ON mv_s.position_id = pos.id " +
		"WHERE mv_s.movie_id = $1 ORDER BY pos.id"
	sqlGetByPersonID = "SELECT p.id, p.name, p.photo, p.addit_photo1, p.addit_photo2, p.description FROM person AS p " +
		"WHERE p.id = $1"
)

func (ss *staffStorage) GetByMovieID(id int) ([]*proto.PersonInMovie, error) {
	movieStaff := make([]*proto.PersonInMovie, 0)

	rows, err := ss.db.Query(sqlGetByMovieID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nextPerson proto.PersonInMovie
		if err = rows.Scan(&nextPerson.ID, &nextPerson.Name, &nextPerson.Photo, &nextPerson.Position); err != nil {
			return nil, err
		}
		movieStaff = append(movieStaff, &nextPerson)
	}

	return movieStaff, nil
}

func (ss *staffStorage) GetByPersonID(id int) (*internal.Person, error) {
	var selectedPerson internal.Person

	err := ss.db.QueryRow(sqlGetByPersonID, id).Scan(&selectedPerson.ID, &selectedPerson.Name,
		&selectedPerson.Photo, &selectedPerson.AdditPhoto1, &selectedPerson.AdditPhoto2, &selectedPerson.Description)

	if err != nil {
		return nil, err
	}

	return &selectedPerson, nil
}
