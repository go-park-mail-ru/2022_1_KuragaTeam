package movie

import (
	"myapp/internal"
)

type Service interface {
	GetByID(id int) (*internal.Movie, error)
	GetRandom(limit, offset int) ([]internal.Movie, error)
	GetMainMovie() (*internal.MainMovieInfoDTO, error)
}
