package delivery

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	movie "myapp/internal/microservices/movie/proto"
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
	logger *zap.SugaredLogger

	movieMicroservice movie.MoviesClient
}

func NewMovieHandler(client movie.MoviesClient, logger *zap.SugaredLogger) *handler {
	return &handler{movieMicroservice: client, logger: logger}
}

func (h *handler) Register(router *echo.Echo) {
	router.GET(moviesURL, h.GetRandomMovies())
	router.GET(movieURL, h.GetMovie())
	router.GET(mainMovieURL, h.GetMainMovie())
}

func (h *handler) GetMovie() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)

		movieID, err := strconv.Atoi(ctx.Param("movie_id"))
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		selectedMovie, err := h.movieMicroservice.GetByID(context.Background(), &movie.GetMovieOptions{MovieID: int64(movieID)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return ctx.JSON(http.StatusOK, selectedMovie)
	}
}

func (h *handler) GetRandomMovies() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		limitStr := ctx.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = randomCount
		}
		movies, err := h.movieMicroservice.GetRandom(context.Background(), &movie.GetRandomOptions{Limit: int32(limit), Offset: offset})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return ctx.JSON(http.StatusOK, &movies.Movie)
	}
}

func (h *handler) GetMainMovie() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		mainMovie, err := h.movieMicroservice.GetMainMovie(context.Background(), &movie.GetMainMovieOptions{})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			return ctx.JSON(http.StatusInternalServerError, &Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		return ctx.JSON(http.StatusOK, &mainMovie)
	}
}
