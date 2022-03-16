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

func (ss *staffStorage) GetByMovieID(id int) ([]domain.Person, error) {
	movieStaff := make([]domain.Person, 0)

	sql := "SELECT p.id, p.name, p.photo FROM person AS p JOIN movies_staff mv_s ON mv_s.person_id = p.id " +
		"WHERE mv_s.movie_id = $1"
	rows, err := ss.db.Query(context.Background(), sql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var nextPerson domain.Person
		if err = rows.Scan(&nextPerson.ID, &nextPerson.Name, &nextPerson.Photo); err != nil {
			return nil, err
		}
		movieStaff = append(movieStaff, nextPerson)
	}

	return movieStaff, nil
}
