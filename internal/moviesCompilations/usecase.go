package moviesCompilations

import (
	"github.com/labstack/echo/v4"
	"myapp/internal"
)

type Service interface {
	GetMainCompilations(context echo.Context) ([]internal.MovieCompilation, error)
	GetByMovieID(context echo.Context, limit int) (internal.MovieCompilation, error)
}
