package compilations

import "myapp/internal/microservices/compilations/proto"

type Storage interface {
	GetByGenre(genreID int) (*proto.MovieCompilation, error)
	GetByCountry(countryID int) (*proto.MovieCompilation, error)
	GetByMovie(movieID int) (*proto.MovieCompilation, error)
	GetByPerson(personID int) (*proto.MovieCompilation, error)
	GetTop(limit int) (*proto.MovieCompilation, error)
	GetTopByYear(year int) (*proto.MovieCompilation, error)
}
