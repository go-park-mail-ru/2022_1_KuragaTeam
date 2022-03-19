package moviesCompilations

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/domain"
)

type Service interface {
	GetMainCompilations(context echo.Context) ([]domain.MovieCompilation, error)
	GetByMovieID(context echo.Context, limit int) (domain.MovieCompilation, error)
}
