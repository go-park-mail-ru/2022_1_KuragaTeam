package staff

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"myapp/internal/domain"
	"myapp/internal/domain/staff"
)

type staffStorage struct {
	db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) staff.Storage {
	return &staffStorage{db: db}
}

const (
	sqlGetByMovieID = "SELECT p.id, p.name, p.photo, pos.name FROM person AS p JOIN movies_staff mv_s ON mv_s.person_id = p.id " +
		"JOIN position pos ON mv_s.position_id = pos.id " +
		"WHERE mv_s.movie_id = $1 ORDER BY pos.id"
	sqlGetByPersonID = "SELECT p.id, p.name, p.photo, p.description FROM person AS p " +
		"WHERE p.id = $1"
)

func (ss *staffStorage) GetByMovieID(id int) ([]domain.PersonInMovieDTO, error) {
	movieStaff := make([]domain.PersonInMovieDTO, 0)

	rows, err := ss.db.Query(context.Background(), sqlGetByMovieID, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var nextPerson domain.PersonInMovieDTO
		if err = rows.Scan(&nextPerson.ID, &nextPerson.Name, &nextPerson.Photo, &nextPerson.Position); err != nil {
			return nil, err
		}
		movieStaff = append(movieStaff, nextPerson)
	}

	return movieStaff, nil
}

func (ss *staffStorage) GetByPersonID(id int) (*domain.Person, error) {
	var selectedPerson domain.Person

	err := ss.db.QueryRow(context.Background(), sqlGetByPersonID, id).Scan(&selectedPerson.ID, &selectedPerson.Name,
		&selectedPerson.Photo, &selectedPerson.Description)
	if err != nil {
		return nil, err
	}

	return &selectedPerson, nil
}