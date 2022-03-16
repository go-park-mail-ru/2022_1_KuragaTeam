package movie

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/adapters/api"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	movieURL     = "api/v1/movie/:movie_id"
	moviesURL    = "api/v1/movies"
	mainMovieURL = "api/v1/mainMovie"
	randomCount  = 10
	offset       = 0
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
	router.GET(mainMovieURL, h.GetMainMovie())
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
		return context.JSON(http.StatusOK, &selectedMovie)
	}
}

func (h *handler) GetRandomMovies() echo.HandlerFunc {
	return func(context echo.Context) error {
		movies, err := h.movieService.GetRandom(randomCount, offset)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &movies)
	}
}

func (h *handler) GetMainMovie() echo.HandlerFunc {
	return func(context echo.Context) error {
		mainMovie, err := h.movieService.GetMainMovie()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusBadRequest, &mainMovie)
	}
}
