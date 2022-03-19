package staff

import "myapp/internal/domain"

type Storage interface {
	GetByMovieID(id int) ([]domain.Person, error)
	GetByPersonID(id int) (*domain.Person, error)
}
