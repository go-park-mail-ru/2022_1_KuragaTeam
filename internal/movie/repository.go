package movie

import (
	"myapp/internal"
)

type Storage interface {
	GetOne(id int) (*internal.Movie, error)
	GetAllMovies(limit, offset int) ([]internal.Movie, error)
	GetRandomMovie() (*internal.MainMovieInfoDTO, error)
}
