package staff

import "myapp/internal/domain"

type Storage interface {
	GetByMovieID(id int) ([]domain.PersonInMovieDTO, error)
	GetByPersonID(id int) (*domain.Person, error)
}
