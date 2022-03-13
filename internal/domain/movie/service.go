package movie

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api/movie"
	"myapp/internal/domain"
)

type service struct {
	storage Storage
}

func NewService(storage Storage) movie.Service {
	return &service{storage: storage}
}

func (s *service) GetByID(context echo.Context, id int) (*domain.Movie, error) {
	return nil, nil
}
func (s *service) GetRandom(context echo.Context, limit int) ([]domain.Movie, error) {
	return s.storage.GetRandom(limit)
}
