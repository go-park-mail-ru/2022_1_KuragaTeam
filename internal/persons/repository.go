package persons

import (
	"myapp/internal"
	compilations "myapp/internal/microservices/compilations/proto"
	movie "myapp/internal/microservices/movie/proto"
)

type Storage interface {
	GetByMovieID(id int) ([]*movie.PersonInMovie, error)
	GetByPersonID(id int) (*internal.Person, error)
	FindPerson(text string) (*compilations.PersonCompilation, error)
	FindPersonByPartial(text string) (*compilations.PersonCompilation, error)
}
