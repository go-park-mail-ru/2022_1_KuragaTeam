package delivery

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"myapp/internal"
	movie "myapp/internal/microservices/movie/proto"
	"net/http"
	"strconv"
)

//type Response struct {
//	Status  int    `json:"status"`
//	Message string `json:"message"`
//}

const (
	movieURL      = "api/v1/movie/:movie_id"
	moviesURL     = "api/v1/oldMovies"
	mainMovieURL  = "api/v1/mainMovie"
	randomCount   = 10
	defaultOffset = 0
)

type handler struct {
	logger *zap.SugaredLogger

	movieMicroservice movie.MoviesClient
}

func mapMovie(inputMovie *movie.Movie) *internal.Movie {
	mappedMovie := &internal.Movie{
		ID:              int(inputMovie.ID),
		Name:            inputMovie.Name,
		IsMovie:         inputMovie.IsMovie,
		NamePicture:     inputMovie.NamePicture,
		Year:            int(inputMovie.Year),
		Duration:        inputMovie.Duration,
		AgeLimit:        int(inputMovie.AgeLimit),
		Description:     inputMovie.Description,
		KinopoiskRating: inputMovie.KinopoiskRating,
		Rating:          inputMovie.Rating,
		Tagline:         inputMovie.Tagline,
		Picture:         inputMovie.Picture,
		Video:           inputMovie.Video,
		Trailer:         inputMovie.Trailer,
		Country:         inputMovie.Country,
	}
	for _, person := range inputMovie.Staff {
		newPerson := internal.PersonInMovieDTO{
			ID:       int(person.ID),
			Name:     person.Name,
			Photo:    person.Photo,
			Position: person.Position,
		}
		mappedMovie.Staff = append(mappedMovie.Staff, newPerson)
	}

	for _, genre := range inputMovie.Genre {
		newGenre := internal.Genre{
			ID:   int(genre.ID),
			Name: genre.Name,
		}
		mappedMovie.Genre = append(mappedMovie.Genre, newGenre)
	}

	if !inputMovie.IsMovie {
		for _, season := range inputMovie.Seasons {
			mappedSeason := internal.Season{
				ID:     int(season.ID),
				Number: int(season.Number),
			}
			for _, episode := range season.Episodes {
				mappedSeason.Episodes = append(mappedSeason.Episodes, internal.Episode{
					ID:          int(episode.ID),
					Name:        episode.Name,
					Number:      int(episode.Number),
					Description: episode.Description,
					Video:       episode.Video,
					Picture:     episode.Picture,
				})
			}
			mappedMovie.Season = append(mappedMovie.Season, mappedSeason)
		}
	}
	return mappedMovie
}
func mapMainMovie(inputMovie *movie.MainMovie) *internal.MainMovieInfoDTO {
	return &internal.MainMovieInfoDTO{
		ID:          int(inputMovie.ID),
		NamePicture: inputMovie.NamePicture,
		Tagline:     inputMovie.Tagline,
		Picture:     inputMovie.Picture,
	}
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

		return ctx.JSON(http.StatusOK, mapMovie(selectedMovie))
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

		offsetStr := ctx.QueryParam("offset")
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			offset = defaultOffset
		}

		movies, err := h.movieMicroservice.GetRandom(context.Background(), &movie.GetRandomOptions{Limit: int32(limit), Offset: int32(offset)})
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
		return ctx.JSON(http.StatusOK, mapMainMovie(mainMovie))
	}
}
