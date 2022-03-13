package movie

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api"
	"myapp/internal/domain"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseMovieRandom struct {
	Status int            `json:"status"`
	Movies []domain.Movie `json:"movies"`
}

const (
	movieURL  = "api/v1/movie/:movie_id"
	moviesURL = "api/v1/movies"
)

type handler struct {
	movieService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{movieService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(moviesURL, h.GetRandomMovies())
}

func (h *handler) GetRandomMovies() echo.HandlerFunc {
	return func(context echo.Context) error {
		movies, err := h.movieService.GetRandom(context, 10)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &ResponseMovieRandom{
			Status: http.StatusOK,
			Movies: movies,
		})
	}
}
