package delivery

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"myapp/internal/moviesCompilations"
	"myapp/mock"
	"net/http/httptest"
	"testing"
)

func TestMoviesCompilationsDelivery_GetMoviesCompilations(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}
	movie3 := moviesCompilations.Movie{
		ID:      3,
		Name:    "Movie3",
		Genre:   []string{"Детектив", "Триллер"},
		Picture: "picture_name3.webp",
	}

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1, movie2},
	}
	MC2 := moviesCompilations.MovieCompilation{
		Name:   "Test MC2",
		Movies: []moviesCompilations.Movie{movie2, movie3},
	}
	MC3 := moviesCompilations.MovieCompilation{
		Name:   "Test MC3",
		Movies: []moviesCompilations.Movie{movie3, movie1},
	}

	tests := []struct {
		name          string
		useCaseMock   *mock.MockMovieCompilationService
		expected      []moviesCompilations.MovieCompilation
		expectedError bool
	}{
		{
			name: "Get main movie",
			useCaseMock: &mock.MockMovieCompilationService{
				GetMainCompilationsFunc: func() ([]moviesCompilations.MovieCompilation, error) {
					return []moviesCompilations.MovieCompilation{MC1, MC2, MC3}, nil
				},
			},
			expected: []moviesCompilations.MovieCompilation{
				MC1, MC2, MC3,
			},
			expectedError: false,
		},
		{
			name: "Return error",
			useCaseMock: &mock.MockMovieCompilationService{
				GetMainCompilationsFunc: func() ([]moviesCompilations.MovieCompilation, error) {
					return nil, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)

			r := NewHandler(test.useCaseMock)
			r.Register(server)
			mainMC := r.GetMoviesCompilations()

			_ = mainMC(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result []moviesCompilations.MovieCompilation
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

func TestMoviesCompilationsDelivery_GetMCByMovieID(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1, movie2},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockMovieCompilationService
		expected      moviesCompilations.MovieCompilation
		expectedError bool
	}{
		{
			name:        "Get movie by ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieCompilationService{
				GetByMovieFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expected:      MC1,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieCompilationService{
				GetByMovieFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:        "No param",
			paramExists: false,
			useCaseMock: &mock.MockMovieCompilationService{
				GetByMovieFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/movie/1", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			if test.paramExists {
				ctx.SetParamNames("movie_id")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock)
			r.Register(server)
			MCByMovieID := r.GetMCByMovieID()

			_ = MCByMovieID(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result moviesCompilations.MovieCompilation
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

func TestMoviesCompilationsDelivery_GetMCByGenreID(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1, movie2},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockMovieCompilationService
		expected      moviesCompilations.MovieCompilation
		expectedError bool
	}{
		{
			name:        "Get MC by genre ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieCompilationService{
				GetByGenreFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expected:      MC1,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieCompilationService{
				GetByGenreFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:        "No param",
			paramExists: false,
			useCaseMock: &mock.MockMovieCompilationService{
				GetByGenreFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/genre/1", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			if test.paramExists {
				ctx.SetParamNames("genre_id")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock)
			r.Register(server)
			MCByMovieID := r.GetMCByGenre()

			_ = MCByMovieID(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result moviesCompilations.MovieCompilation
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

func TestMoviesCompilationsDelivery_GetMCByPersonID(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1, movie2},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockMovieCompilationService
		expected      moviesCompilations.MovieCompilation
		expectedError bool
	}{
		{
			name:        "Get MC by country ID",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieCompilationService{
				GetByPersonFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expected:      MC1,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "1",
			useCaseMock: &mock.MockMovieCompilationService{
				GetByPersonFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:        "No param",
			paramExists: false,
			useCaseMock: &mock.MockMovieCompilationService{
				GetByPersonFunc: func(id int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/person/1", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			if test.paramExists {
				ctx.SetParamNames("person_id")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock)
			r.Register(server)
			MCByMovieID := r.GetMCByPersonID()

			_ = MCByMovieID(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result moviesCompilations.MovieCompilation
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

func TestMoviesCompilationsDelivery_GetTopMC(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1, movie2},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockMovieCompilationService
		expected      moviesCompilations.MovieCompilation
		expectedError bool
	}{
		{
			name:        "Get top MC",
			paramExists: true,
			param:       "15",
			useCaseMock: &mock.MockMovieCompilationService{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expected:      MC1,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "10",
			useCaseMock: &mock.MockMovieCompilationService{
				GetTopFunc: func(limit int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/top", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)

			if test.paramExists {
				ctx.QueryParams().Set("limit", test.param)
			}

			r := NewHandler(test.useCaseMock)
			r.Register(server)
			mainMC := r.GetTopMC()

			_ = mainMC(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result moviesCompilations.MovieCompilation
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

func TestMoviesCompilationsDelivery_GetTopByYear(t *testing.T) {
	//config := zap.NewDevelopmentConfig()
	//config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//prLogger, _ := config.Build()
	//logger := prLogger.Sugar()
	//defer prLogger.Sync()
	const testError = "test error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movie1 := moviesCompilations.Movie{
		ID:      1,
		Name:    "Movie1",
		Genre:   []string{"Боевик", "Триллер"},
		Picture: "picture_name.webp",
	}
	movie2 := moviesCompilations.Movie{
		ID:      2,
		Name:    "Movie2",
		Genre:   []string{"Фантастика", "Семейный"},
		Picture: "picture_name2.webp",
	}

	MC1 := moviesCompilations.MovieCompilation{
		Name:   "Test MC1",
		Movies: []moviesCompilations.Movie{movie1, movie2},
	}

	tests := []struct {
		name          string
		paramExists   bool
		param         string
		useCaseMock   *mock.MockMovieCompilationService
		expected      moviesCompilations.MovieCompilation
		expectedError bool
	}{
		{
			name:        "Get Top by year",
			paramExists: true,
			param:       "2011",
			useCaseMock: &mock.MockMovieCompilationService{
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expected:      MC1,
			expectedError: false,
		},
		{
			name:        "Return error",
			paramExists: true,
			param:       "2011",
			useCaseMock: &mock.MockMovieCompilationService{
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return moviesCompilations.MovieCompilation{}, errors.New(testError)
				},
			},
			expectedError: true,
		},
		{
			name:        "No param",
			paramExists: false,
			useCaseMock: &mock.MockMovieCompilationService{
				GetTopByYearFunc: func(year int) (moviesCompilations.MovieCompilation, error) {
					return MC1, nil
				},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/yearTop/2011", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			if test.paramExists {
				ctx.SetParamNames("year")
				ctx.SetParamValues(test.param)
			}

			r := NewHandler(test.useCaseMock)
			r.Register(server)
			MCByYear := r.GetYearTopMC()

			_ = MCByYear(ctx)
			body := rec.Body
			status := rec.Result().Status
			var result moviesCompilations.MovieCompilation
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
