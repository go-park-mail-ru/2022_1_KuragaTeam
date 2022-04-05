package delivery

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/internal"
	"myapp/internal/mock"
	"net/http/httptest"
	"testing"
)

func TestMovieDelivery_GetMainMovie(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := internal.MainMovieInfoDTO{
		ID:          0,
		NamePicture: "name_picture.webp",
		Tagline:     "This is test movie",
		Picture:     "movie_picture.webp",
	}

	tests := []struct {
		name          string
		useCaseMock   *mock.MockMovieService
		expected      internal.MainMovieInfoDTO
		expectedError bool
	}{
		{
			name: "Get main movie",
			useCaseMock: &mock.MockMovieService{
				GetMainMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
					return &movie, nil
				},
			},
			expected:      movie,
			expectedError: false,
		},
		{
			name: "Return error",
			useCaseMock: &mock.MockMovieService{
				GetMainMovieFunc: func() (*internal.MainMovieInfoDTO, error) {
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

			r := NewHandler(test.useCaseMock)
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
				assert.Equal(t, test.expected, result)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}

func TestMovieDelivery_GetMovie(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie := internal.Movie{
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
		Genre:           []string{"Комедия", "История"},
		Staff: []internal.PersonInMovieDTO{
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
		useCaseMock   *mock.MockMovieService
		expected      internal.Movie
		expectedError bool
	}{
		{
			name:        "Get movie by ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieService{
				GetByIDFunc: func(id int) (*internal.Movie, error) {
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
			useCaseMock: &mock.MockMovieService{
				GetByIDFunc: func(id int) (*internal.Movie, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:        "No param",
			paramExists: false,
			useCaseMock: &mock.MockMovieService{
				GetByIDFunc: func(id int) (*internal.Movie, error) {
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
			if test.paramExists {
				ctx.SetParamNames("movie_id")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock)
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
				assert.Equal(t, test.expected, result)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}

func TestMovieDelivery_GetRandomMovies(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movies := []internal.Movie{{
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
		Genre:           []string{"Комедия", "История"},
		Staff: []internal.PersonInMovieDTO{
			{
				ID:       1,
				Name:     "Актер1",
				Photo:    "actor_1.webp",
				Position: "Актер",
			},
		},
	},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockMovieService
		expected      []internal.Movie
		expectedError bool
	}{
		{
			name:        "Get movie by ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieService{
				GetRandomFunc: func(limit, offset int) ([]internal.Movie, error) {
					return movies, nil
				},
			},
			expected:      movies,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieService{
				GetRandomFunc: func(limit, offset int) ([]internal.Movie, error) {
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
			if test.paramExists {
				ctx.SetParamNames("movie_id")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock)
			//r.Register(server)
			movieByID := r.GetRandomMovies()

			_ = movieByID(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result []internal.Movie
			_ = json.Unmarshal(body.Bytes(), &result)
			if test.expectedError {
				assert.Equal(t, "500 Internal Server Error", status)
			} else {
				assert.Equal(t, test.expected, result)
				assert.Equal(t, "200 OK", status)
			}
		})
	}
}
