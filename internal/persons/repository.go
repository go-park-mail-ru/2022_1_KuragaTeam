package persons

import (
	"myapp/internal"
	"myapp/internal/microservices/movie/proto"
)

type Storage interface {
	GetByMovieID(id int) ([]*proto.PersonInMovie, error)
	GetByPersonID(id int) (*internal.Person, error)
}
