package delivery

import (
	"encoding/json"
	"errors"
	"myapp/internal"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/microservices/movie/usecase"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestMovieDelivery_GetMainMovie(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := proto.MainMovie{
		ID:          0,
		NamePicture: "name_picture.webp",
		Tagline:     "This is test movie",
		Picture:     "movie_picture.webp",
	}

	tests := []struct {
		name          string
		useCaseMock   *usecase.MockMoviesClient
		expected      proto.MainMovie
		expectedError bool
	}{
		{
			name: "Get main movie",
			useCaseMock: &usecase.MockMoviesClient{
				GetMainMovieFunc: func(ctx context.Context, in *proto.GetMainMovieOptions, opts ...grpc.CallOption) (*proto.MainMovie, error) {
					return &movie, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
		{
			name: "Return error",
			useCaseMock: &usecase.MockMoviesClient{
				GetMainMovieFunc: func(ctx context.Context, in *proto.GetMainMovieOptions, opts ...grpc.CallOption) (*proto.MainMovie, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/mainMovie", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")

			r := NewMovieHandler(test.useCaseMock, logger)
			r.Register(server)
			mainMovie := r.GetMainMovie()

			_ = mainMovie(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result internal.MainMovieInfoDTO
			_ = json.Unmarshal(body.Bytes(), &result)
			if test.expectedError {
				assert.Equal(t, "500 Internal Server Error", status)
			} else {
				assert.Equal(t, int(test.expected.ID), result.ID)
				assert.Equal(t, test.expected.NamePicture, result.NamePicture)
				assert.Equal(t, test.expected.Tagline, result.Tagline)
				assert.Equal(t, test.expected.Picture, result.Picture)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}

func TestMovieDelivery_GetMovie(t *testing.T) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := proto.Movie{
		ID:              1,
		Name:            "Movie1",
		NamePicture:     "name_picture.webp",
		Year:            2010,
		Duration:        "1 час 15 минут",
		AgeLimit:        18,
		Description:     "Это описание тестового фильма",
		KinopoiskRating: 7.5,
		Rating:          9.1,
		Tagline:         "This is test movie",
		Picture:         "movie_picture.webp",
		Video:           "test_movie_video.webm",
		Trailer:         "test_movie_genre.webm",
		Country:         []string{"Россия", "Армения"},
		Genre: []*proto.Genres{
			{
				ID:   1,
				Name: "Комедия",
			},
			{
				ID:   2,
				Name: "История",
			},
		},
		Staff: []*proto.PersonInMovie{
			{
				ID:       1,
				Name:     "Актер1",
				Photo:    "actor_1.webp",
				Position: "Актер",
			},
		},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *usecase.MockMoviesClient
		expected      proto.Movie
		expectedError bool
	}{
		{
			name:        "Get movie by ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &usecase.MockMoviesClient{
				GetByIDFunc: func(ctx context.Context, in *proto.GetMovieOptions, opts ...grpc.CallOption) (*proto.Movie, error) {
					return &movie, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "1",
			useCaseMock: &usecase.MockMoviesClient{
				GetByIDFunc: func(ctx context.Context, in *proto.GetMovieOptions, opts ...grpc.CallOption) (*proto.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:        "No param",
			paramExists: false,
			useCaseMock: &usecase.MockMoviesClient{
				GetByIDFunc: func(ctx context.Context, in *proto.GetMovieOptions, opts ...grpc.CallOption) (*proto.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movie/1", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set("REQUEST_ID", "1")
			if test.paramExists {
				ctx.SetParamNames("movie_id")
				ctx.SetParamValues(test.param)
			}

			r := NewMovieHandler(test.useCaseMock, logger)
			//r.Register(server)
			movieByID := r.GetMovie()

			_ = movieByID(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result internal.Movie
			_ = json.Unmarshal(body.Bytes(), &result)
			if test.expectedError {
				assert.Equal(t, "500 Internal Server Error", status)
			} else {
				assert.Equal(t, test.expected.Name, result.Name)
				assert.Equal(t, test.expected.ID, int64(result.ID))
				assert.Equal(t, test.expected.NamePicture, result.NamePicture)
				assert.Equal(t, test.expected.IsMovie, result.IsMovie)
				assert.Equal(t, test.expected.Picture, result.Picture)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}

//func TestMovieDelivery_GetRandomMovies(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movies := []internal.Movie{{
//		ID:              1,
//		Name:            "Movie1",
//		NamePicture:     "name_picture.webp",
//		Year:            2010,
//		Duration:        "1 час 15 минут",
//		AgeLimit:        18,
//		Description:     "Это описание тестового фильма",
//		KinopoiskRating: 7.5,
//		Rating:          9.1,
//		Tagline:         "This is test movie",
//		Picture:         "movie_picture.webp",
//		Video:           "test_movie_video.webm",
//		Trailer:         "test_movie_genre.webm",
//		Country:         []string{"Россия", "Армения"},
//		Genre:           []string{"Комедия", "История"},
//		Staff: []internal.PersonInMovieDTO{
//			{
//				ID:       1,
//				Name:     "Актер1",
//				Photo:    "actor_1.webp",
//				Position: "Актер",
//			},
//		},
//	},
//	}
//
//	tests := []struct {
//		name          string
//		paramExists   bool
//		param         string
//		useCaseMock   *usecase.MockMoviesClient
//		expected      []internal.Movie
//		expectedError bool
//	}{
//		{
//			name:        "Get movie by ID",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &usecase.MockMoviesClient{
//				GetRandomFunc: func(limit, offset int) ([]internal.Movie, error) {
//					return movies, nil
//				},
//			},
//			expected:      movies,
//			expectedError: false,
//		},
//		{
//			name:        "Return error",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &usecase.MockMoviesClient{
//				GetRandomFunc: func(limit, offset int) ([]internal.Movie, error) {
//					return nil, errors.New(testError)
//				},
//			},
//			expectedError: true,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			server := echo.New()
//			req := httptest.NewRequest(echo.GET, "/api/v1/movie/1", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//			if test.paramExists {
//				ctx.SetParamNames("movie_id")
//				ctx.SetParamValues(test.param)
//			}
//
//			r := NewHandler(test.useCaseMock, logger)
//			//r.Register(server)
//			movieByID := r.GetRandomMovies()
//
//			_ = movieByID(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result []internal.Movie
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				assert.Equal(t, test.expected, result)
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
