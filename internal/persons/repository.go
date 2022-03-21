package persons

import (
	"myapp/internal"
)

type Storage interface {
	GetByMovieID(id int) ([]internal.PersonInMovieDTO, error)
	GetByPersonID(id int) (*internal.Person, error)
}
