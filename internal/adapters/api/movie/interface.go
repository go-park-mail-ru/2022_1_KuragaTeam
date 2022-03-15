package movie

import (
	"myapp/internal/domain"
)

type Service interface {
	GetByID(id int) (*domain.Movie, error)
	GetRandom(limit int) ([]domain.Movie, error)
}
