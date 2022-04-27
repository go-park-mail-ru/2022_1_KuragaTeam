package delivery

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	compilations "myapp/internal/microservices/compilations/proto"
	"net/http"
	"strconv"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	MCAllMoviesURL = "api/v1/movies"
	MCAllSeriesURL = "api/v1/series"
	MCByPersonURL  = "api/v1/movieCompilations/person/:person_id"
	MCByMovieURL   = "api/v1/movieCompilations/movie/:movie_id"
	MCByGenreURL   = "api/v1/movieCompilations/genre/:genre_id"
	MCTopURL       = "api/v1/movieCompilations/top"
	MCYearTopURL   = "api/v1/movieCompilations/yearTop/:year"
	MCDefaultURL   = "/api/v1/movieCompilations"
)

type compilationsHandler struct {
	compilationsMicroservice compilations.MovieCompilationsClient
	logger                   *zap.SugaredLogger
}

func NewHandler(service compilations.MovieCompilationsClient, logger *zap.SugaredLogger) *compilationsHandler {
	return &compilationsHandler{compilationsMicroservice: service, logger: logger}
}

func (h *compilationsHandler) Register(router *echo.Echo) {
	router.GET(MCAllMoviesURL, h.GetAllMovies())
	router.GET(MCAllSeriesURL, h.GetAllSeries())
	router.GET(MCDefaultURL, h.GetMoviesCompilations())
	router.GET(MCByMovieURL, h.GetMCByMovieID())
	router.GET(MCByGenreURL, h.GetMCByGenre())
	router.GET(MCByPersonURL, h.GetMCByPersonID())
	router.GET(MCTopURL, h.GetTopMC())
	router.GET(MCYearTopURL, h.GetYearTopMC())
}

func (h *compilationsHandler) GetAllMovies() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		limitStr := ctx.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = randomCount
		}

		offsetStr := ctx.QueryParam("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = defaultOffset
		}

		selectedMC, err := h.compilationsMicroservice.GetAllMovies(context.Background(), &compilations.GetCompilationOptions{Limit: int32(limit), Offset: int32(offset)})
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
		return ctx.JSON(http.StatusOK, selectedMC.Movies)
	}
}

func (h *compilationsHandler) GetAllSeries() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		limitStr := ctx.QueryParam("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = randomCount
		}

		offsetStr := ctx.QueryParam("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = defaultOffset
		}

		selectedMC, err := h.compilationsMicroservice.GetAllSeries(context.Background(), &compilations.GetCompilationOptions{Limit: int32(limit), Offset: int32(offset)})
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
		return ctx.JSON(http.StatusOK, selectedMC.Movies)
	}
}

func (h *compilationsHandler) GetMoviesCompilations() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		mainMoviesCompilations, err := h.compilationsMicroservice.GetMainCompilations(context.Background(), &compilations.GetMainCompilationsOptions{})
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
		return ctx.JSON(http.StatusOK, mainMoviesCompilations)
	}
}

func (h *compilationsHandler) GetMCByMovieID() echo.HandlerFunc {
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
		selectedMC, err := h.compilationsMicroservice.GetByMovie(context.Background(), &compilations.GetByIDOptions{ID: int32(movieID)})
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
		return ctx.JSON(http.StatusOK, selectedMC)
	}
}

func (h *compilationsHandler) GetMCByGenre() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		genreID, err := strconv.Atoi(ctx.Param("genre_id"))
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
		selectedMC, err := h.compilationsMicroservice.GetByGenre(context.Background(), &compilations.GetByIDOptions{ID: int32(genreID)})
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
		return ctx.JSON(http.StatusOK, selectedMC)
	}
}

func (h *compilationsHandler) GetMCByPersonID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		personID, err := strconv.Atoi(ctx.Param("person_id"))
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
		selectedMC, err := h.compilationsMicroservice.GetByPerson(context.Background(), &compilations.GetByIDOptions{ID: int32(personID)})
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
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return ctx.JSON(http.StatusOK, selectedMC)
	}
}

func (h *compilationsHandler) GetTopMC() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		var limit int
		echo.QueryParamsBinder(ctx).Int("limit", &limit)
		if limit > 12 {
			limit = 12
		}
		selectedMC, err := h.compilationsMicroservice.GetTop(context.Background(), &compilations.GetCompilationOptions{Limit: int32(limit)})
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
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return ctx.JSON(http.StatusOK, selectedMC)
	}
}

func (h *compilationsHandler) GetYearTopMC() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID := ctx.Get("REQUEST_ID").(string)
		year, err := strconv.Atoi(ctx.Param("year"))
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
		selectedMC, err := h.compilationsMicroservice.GetTopByYear(context.Background(), &compilations.GetByIDOptions{ID: int32(year)})
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
			zap.Int("ANSWER STATUS", http.StatusInternalServerError),
		)
		return ctx.JSON(http.StatusOK, selectedMC)
	}
}
