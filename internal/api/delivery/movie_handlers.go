package delivery

import (
	"github.com/mailru/easyjson"
	"myapp/internal"
	"myapp/internal/constants"
	movie "myapp/internal/microservices/movie/proto"
	"myapp/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

//type Response struct {
//	Status  int    `json:"status"`
//	Message string `json:"message"`
//}

const (
	movieURL      = "api/v1/movie/:movie_id"
	mainMovieURL  = "api/v1/mainMovie"
	addRatingURL  = "api/v1/addMovieRating"
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
	router.GET(movieURL, h.GetMovie())
	router.GET(mainMovieURL, h.GetMainMovie())
	router.POST(addRatingURL, h.AddMovieRating())
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
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		selectedMovie, err := h.movieMicroservice.GetByID(context.Background(), &movie.GetMovieOptions{MovieID: int64(movieID)})
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)

		resp, err := easyjson.Marshal(mapMovie(selectedMovie))
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
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
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}
		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
		)
		resp, err := easyjson.Marshal(mapMainMovie(mainMovie))
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}
		return ctx.JSONBlob(http.StatusOK, resp)
	}
}

func (h *handler) AddMovieRating() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		requestID, ok := ctx.Get("REQUEST_ID").(string)
		if !ok {
			h.logger.Error(
				zap.String("ERROR", constants.NoRequestId),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError))
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: constants.NoRequestId,
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		userID, ok := ctx.Get("USER_ID").(int64)
		if !ok {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.SessionRequired),
				zap.Int("ANSWER STATUS", http.StatusBadRequest),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusBadRequest,
				Message: constants.SessionRequired,
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusBadRequest, resp)
		}
		if userID == -1 {
			h.logger.Info(
				zap.String("ID", requestID),
				zap.String("ERROR", constants.UserIsUnauthorized),
				zap.Int("ANSWER STATUS", http.StatusUnauthorized),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusUnauthorized,
				Message: constants.UserIsUnauthorized,
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusUnauthorized, resp)
		}

		requestOptions := internal.MovieRatingDTO{}

		if err := ctx.Bind(&requestOptions); err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		data := &movie.AddRatingOptions{
			UserID:  userID,
			MovieID: int64(requestOptions.MovieID),
			Rating:  int32(requestOptions.Rating),
		}

		newRating, err := h.movieMicroservice.AddMovieRating(context.Background(), data)
		if err != nil {
			h.logger.Error(
				zap.String("ID", requestID),
				zap.String("ERROR", err.Error()),
				zap.Int("ANSWER STATUS", http.StatusInternalServerError),
			)
			resp, err := easyjson.Marshal(&models.Response{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
			if err != nil {
				return ctx.NoContent(http.StatusInternalServerError)
			}
			return ctx.JSONBlob(http.StatusInternalServerError, resp)
		}

		h.logger.Info(
			zap.String("ID", requestID),
			zap.Int("ANSWER STATUS", http.StatusOK),
			zap.Float32("NEW RATING: ", newRating.Rating),
		)

		return ctx.JSON(http.StatusOK, &newRating.Rating)
	}
}
