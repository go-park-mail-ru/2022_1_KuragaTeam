package movie

import (
	"myapp/internal/domain"
)

type Service interface {
	GetByID(id int) (*domain.Movie, error)
	GetRandom(limit, offset int) ([]domain.Movie, error)
	GetMainMovie() (*domain.MainMovieInfoDTO, error)
}
