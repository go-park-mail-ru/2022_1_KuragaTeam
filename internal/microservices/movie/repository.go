package movie

import (
	"myapp/internal/microservices/movie/proto"
)

type Storage interface {
	GetOne(id int) (*proto.Movie, error)
	GetAllMovies(limit, offset int) ([]*proto.Movie, error)
	GetSeasonsAndEpisodes(seriesId int) ([]*proto.Season, error)
	GetRandomMovie() (*proto.MainMovie, error)
}
