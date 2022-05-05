package repository

import (
	"database/sql"
	"myapp/internal"
	"myapp/internal/constants"
	compilations "myapp/internal/microservices/compilations/proto"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/persons"

	"github.com/lib/pq"
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
	findPerson = "select p.id, p.name, p.photo," +
		"array((select distinct position.name from position join movies_staff ms on" +
		" position.id = ms.position_id where ms.person_id = p.id)) from person as p " +
		"where to_tsvector('russian', p.name) @@ to_tsquery('russian', $1) ORDER BY p.name LIMIT $2;"
	findPersonByPartial = "select p.id, p.name, p.photo," +
		"array((select distinct position.name from position join movies_staff ms on" +
		" position.id = ms.position_id where ms.person_id = p.id)) from person as p " +
		"where p.name ILIKE $1 ORDER BY p.name LIMIT $2;"
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

func (ss *staffStorage) FindPerson(text string) (*compilations.PersonCompilation, error) {
	var selectedPersonCompilation compilations.PersonCompilation

	rows, err := ss.db.Query(findPerson, text, constants.PersonsSearchLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedPersons compilations.PersonInfo
		err = rows.Scan(&selectedPersons.ID, &selectedPersons.Name, &selectedPersons.Photo, pq.Array(&selectedPersons.Position))
		if err != nil {
			return nil, err
		}
		selectedPersonCompilation.Persons = append(selectedPersonCompilation.Persons, &selectedPersons)
	}
	return &selectedPersonCompilation, nil
}

func (ss *staffStorage) FindPersonByPartial(text string) (*compilations.PersonCompilation, error) {
	var selectedPersonCompilation compilations.PersonCompilation

	rows, err := ss.db.Query(findPersonByPartial, "%"+text+"%", constants.PersonsSearchLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var selectedPersons compilations.PersonInfo
		err = rows.Scan(&selectedPersons.ID, &selectedPersons.Name, &selectedPersons.Photo, pq.Array(&selectedPersons.Position))
		if err != nil {
			return nil, err
		}
		selectedPersonCompilation.Persons = append(selectedPersonCompilation.Persons, &selectedPersons)
	}
	return &selectedPersonCompilation, nil
}
