package delivery

import (
	"github.com/labstack/echo/v4"
	"myapp/internal"
	"myapp/internal/api"
	"myapp/internal/moviesCompilations"
	"net/http"
)

type ResponseMovieCompilations struct {
	Status           int                         `json:"status"`
	MovieCompilation []internal.MovieCompilation `json:"moviesCompilation"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	MCByMovieIDURL = "api/v1/movieCompilations/:movie_id"
	MCDefaultURL   = "/api/v1/movieCompilations"
)

type handler struct {
	movieCompilationsService moviesCompilations.Service
}

func NewHandler(service moviesCompilations.Service) api.Handler {
	return &handler{movieCompilationsService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(MCDefaultURL, h.GetMoviesCompilations())
	router.GET(MCByMovieIDURL, h.GetMCByMovieID())
}

func (h *handler) GetMoviesCompilations() echo.HandlerFunc {
	return func(context echo.Context) error {
		mainMoviesCompilations, err := h.movieCompilationsService.GetMainCompilations(context)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &ResponseMovieCompilations{
			Status:           http.StatusOK,
			MovieCompilation: mainMoviesCompilations,
		})
	}
}
func (h *handler) GetMCByMovieID() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusInternalServerError, &Response{
			Status:  http.StatusInternalServerError,
			Message: "Not realised",
		})
	}
}
