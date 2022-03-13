package movie

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/domain"
)

type Service interface {
	GetByID(context echo.Context, id int) (*domain.Movie, error)
	GetRandom(context echo.Context, limit int) ([]domain.Movie, error)
}
