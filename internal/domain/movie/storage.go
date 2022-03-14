package movie

import "myapp/internal/domain"

type Storage interface {
	GetOne(id int) (*domain.Movie, error)
	GetRandom(limit int) ([]domain.Movie, error)
}
