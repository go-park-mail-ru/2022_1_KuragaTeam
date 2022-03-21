package movie

import (
	"myapp/internal"
)

type Storage interface {
	GetOne(id int) (*internal.Movie, error)
	GetRandomMovies(limit, offset int) ([]internal.Movie, error)
	GetRandomMovie() (*internal.MainMovieInfoDTO, error)
}
