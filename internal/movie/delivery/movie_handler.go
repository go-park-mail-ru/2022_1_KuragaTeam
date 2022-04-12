package delivery

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"myapp/internal/movie"
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
	movieService movie.Service
	logger       *zap.SugaredLogger
}

func NewHandler(service movie.Service, logger *zap.SugaredLogger) *handler {
	return &handler{movieService: service, logger: logger}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(moviesURL, h.GetRandomMovies())
	router.GET(movieURL, h.GetMovie())
	router.GET(mainMovieURL, h.GetMainMovie())
}

func (h *handler) GetMovie() echo.HandlerFunc {
	return func(context echo.Context) error {
		requestID := context.Get("REQUEST_ID").(string)

		movieID, err := strconv.Atoi(context.Param("movie_id"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		selectedMovie, err := h.movieService.GetByID(movieID)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return context.JSON(http.StatusOK, &selectedMovie)
	}
}

func (h *handler) GetRandomMovies() echo.HandlerFunc {
	return func(context echo.Context) error {
		requestID := context.Get("REQUEST_ID").(string)
		movies, err := h.movieService.GetRandom(randomCount, offset)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return context.JSON(http.StatusOK, &movies)
	}
}

func (h *handler) GetMainMovie() echo.HandlerFunc {
	return func(context echo.Context) error {
		requestID := context.Get("REQUEST_ID").(string)
		mainMovie, err := h.movieService.GetMainMovie()
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return context.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return context.JSON(http.StatusOK, &mainMovie)
	}
}
