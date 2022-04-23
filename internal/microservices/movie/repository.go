package movie

import (
	"myapp/internal/microservices/movie/proto"
)

type Storage interface {
	GetOne(id int) (*proto.Movie, error)
	GetAllMovies(limit, offset int) ([]*proto.Movie, error)
	GetRandomMovie() (*proto.MainMovie, error)
}
