package movie

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api"
	"myapp/internal/domain"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseMovieRandom struct {
	Status int            `json:"status"`
	Movies []domain.Movie `json:"movies"`
}
type ResponseMovie struct {
	Status int          `json:"status"`
	Movies domain.Movie `json:"movie"`
}

const (
	movieURL    = "api/v1/movie/:movie_id"
	moviesURL   = "api/v1/movies"
	randomCount = 10
)

type handler struct {
	movieService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{movieService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(moviesURL, h.GetRandomMovies())
	router.GET(movieURL, h.GetMovie())
}

func (h *handler) GetMovie() echo.HandlerFunc {
	return func(context echo.Context) error {
		movieID, err := strconv.Atoi(context.Param("movie_id"))
		selectedMovie, err := h.movieService.GetByID(movieID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &ResponseMovie{
			Status: http.StatusOK,
			Movies: *selectedMovie,
		})
	}
}

func (h *handler) GetRandomMovies() echo.HandlerFunc {
	return func(context echo.Context) error {
		movies, err := h.movieService.GetRandom(randomCount)
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
