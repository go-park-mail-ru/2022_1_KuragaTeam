package moviesCompilations

import (
	"github.com/labstack/echo/v4"
)

type Service interface {
	GetMainCompilations(context echo.Context) ([]MovieCompilation, error)
	GetByGenre(context echo.Context, genreID int) (MovieCompilation, error)
	//GetPopularID(context echo.Context, limit int) (MovieCompilation, error)
}
