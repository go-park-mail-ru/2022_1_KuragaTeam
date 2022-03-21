package delivery

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/api"
	"myapp/internal/moviesCompilations"
	"net/http"
	"strconv"
)

type ResponseMovieCompilations struct {
	Status           int                                   `json:"status"`
	MovieCompilation []moviesCompilations.MovieCompilation `json:"movies_compilation"`
}

type ResponseOneMovieCompilation struct {
	Status           int                                 `json:"status"`
	MovieCompilation moviesCompilations.MovieCompilation `json:"movies_compilation"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	MCByPersonURL = "api/v1/movieCompilations/person/:person_id"
	MCByMovieURL  = "api/v1/movieCompilations/movie/:movie_id"
	MCByGenreURL  = "api/v1/movieCompilations/genre/:genre_id"
	MCDefaultURL  = "/api/v1/movieCompilations"
)

type handler struct {
	movieCompilationsService moviesCompilations.Service
}

func NewHandler(service moviesCompilations.Service) api.Handler {
	return &handler{movieCompilationsService: service}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(MCDefaultURL, h.GetMoviesCompilations())
	router.GET(MCByMovieURL, h.GetMCByMovieID())
	router.GET(MCByGenreURL, h.GetMCByGenre())
}

func (h *handler) GetMoviesCompilations() echo.HandlerFunc {
	return func(context echo.Context) error {
		mainMoviesCompilations, err := h.movieCompilationsService.GetMainCompilations()
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
		movieID, err := strconv.Atoi(context.Param("movie_id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.movieCompilationsService.GetByMovie(movieID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}

func (h *handler) GetMCByGenre() echo.HandlerFunc {
	return func(context echo.Context) error {
		genreID, err := strconv.Atoi(context.Param("genre_id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMC, err := h.movieCompilationsService.GetByGenre(genreID)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		return context.JSON(http.StatusOK, &selectedMC)
	}
}
