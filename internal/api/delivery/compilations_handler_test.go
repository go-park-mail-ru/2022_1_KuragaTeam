package delivery

//
//import (
//	"encoding/json"
//	"errors"
//	"github.com/golang/mock/gomock"
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"myapp/internal/microservices/compilations/proto"
//	mock "myapp/internal/microservices/compilations/usecase"
//	profile "myapp/internal/microservices/profile/usecase"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestMoviesCompilationsDelivery_GetMoviesCompilations(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movie1 := proto.MovieInfo{
//		ID:   1,
//		Name: "Movie1",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Боевик",
//			},
//			{
//				ID:   2,
//				Name: "Криминал",
//			},
//		},
//		Picture: "picture_name.webp",
//	}
//	movie2 := proto.MovieInfo{
//		ID:   2,
//		Name: "Movie2",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Фантастика",
//			},
//			{
//				ID:   2,
//				Name: "Семейный",
//			},
//		},
//		Picture: "picture_name2.webp",
//	}
//	movie3 := proto.MovieInfo{
//		ID:   3,
//		Name: "Movie3",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Детектив",
//			},
//			{
//				ID:   2,
//				Name: "Триллер",
//			},
//		},
//		Picture: "picture_name3.webp",
//	}
//
//	MC1Movie1 := movie1
//	MC1Movie2 := movie2
//	MC1 := proto.MovieCompilation{
//		Name: "Test MC1",
//		Movies: []*proto.MovieInfo{
//			&MC1Movie1,
//			&MC1Movie2,
//		},
//	}
//	MC2Movie2 := movie2
//	MC2Movie3 := movie3
//	MC2 := proto.MovieCompilation{
//		Name: "Test MC2",
//		Movies: []*proto.MovieInfo{
//			&MC2Movie2,
//			&MC2Movie3,
//		},
//	}
//	MC3Movie3 := movie3
//	MC3Movie1 := movie1
//	MC3 := proto.MovieCompilation{
//		Name: "Test MC3",
//		Movies: []*proto.MovieInfo{
//			&MC3Movie3,
//			&MC3Movie1,
//		},
//	}
//
//	tests := []struct {
//		name               string
//		useCaseMock        *mock.MockMovieCompilationsClient
//		profileUsecaseMock *profile.MockProfileClient
//		expected           *proto.MovieCompilationsArr
//		expectedError      bool
//	}{
//		{
//			name: "Get main movie",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetMainCompilationsFunc: func(ctx context.Context, in *proto.GetMainCompilationsOptions, opts ...grpc.CallOption) (*proto.MovieCompilationsArr, error) {
//					return &proto.MovieCompilationsArr{MovieCompilations: []*proto.MovieCompilation{&MC1, &MC2, &MC3}}, nil
//				},
//			},
//			expected: &proto.MovieCompilationsArr{
//				MovieCompilations: []*proto.MovieCompilation{
//					&MC1, &MC2, &MC3,
//				},
//			},
//			expectedError: false,
//		},
//		{
//			name: "Return error",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetMainCompilationsFunc: func(ctx context.Context, in *proto.GetMainCompilationsOptions, opts ...grpc.CallOption) (*proto.MovieCompilationsArr, error) {
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
//			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//
//			r := NewCompilationsHandler(test.useCaseMock, test.profileUsecaseMock, logger)
//			r.Register(server)
//			mainMC := r.GetMoviesCompilations()
//
//			_ = mainMC(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result []proto.MovieCompilation
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				for i, MC := range test.expected.MovieCompilations {
//					for j, movie := range MC.Movies {
//						assert.Equal(t, movie.ID, result[i].Movies[j].ID)
//						assert.Equal(t, movie.Name, result[i].Movies[j].Name)
//						assert.Equal(t, movie.Picture, result[i].Movies[j].Picture)
//						for k, genre := range movie.Genre {
//							assert.Equal(t, genre.ID, result[i].Movies[j].Genre[k].ID)
//							assert.Equal(t, genre.Name, result[i].Movies[j].Genre[k].Name)
//						}
//					}
//				}
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
//
//func TestMoviesCompilationsDelivery_GetMCByMovieID(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movie1 := proto.MovieInfo{
//		ID:   1,
//		Name: "Movie1",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Боевик",
//			},
//			{
//				ID:   2,
//				Name: "Криминал",
//			},
//		},
//		Picture: "picture_name.webp",
//	}
//	movie2 := proto.MovieInfo{
//		ID:   2,
//		Name: "Movie2",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Фантастика",
//			},
//			{
//				ID:   2,
//				Name: "Семейный",
//			},
//		},
//		Picture: "picture_name2.webp",
//	}
//
//	MC1Movie1 := movie1
//	MC1Movie2 := movie2
//	MC1 := proto.MovieCompilation{
//		Name: "Test MC1",
//		Movies: []*proto.MovieInfo{
//			&MC1Movie1,
//			&MC1Movie2,
//		},
//	}
//
//	tests := []struct {
//		name               string
//		paramExists        bool
//		param              string
//		useCaseMock        *mock.MockMovieCompilationsClient
//		profileUsecaseMock *profile.MockProfileClient
//		expected           proto.MovieCompilation
//		expectedError      bool
//	}{
//		{
//			name:        "Get movie by ID",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByMovieFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expected:      MC1,
//			expectedError: false,
//		},
//		{
//			name:        "Return error",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByMovieFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return nil, errors.New(testError)
//				},
//			},
//			expectedError: true,
//		},
//		{
//			name:        "No param",
//			paramExists: false,
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByMovieFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expectedError: true,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			server := echo.New()
//			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/movie/1", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//			if test.paramExists {
//				ctx.SetParamNames("movie_id")
//				ctx.SetParamValues(test.param)
//			}
//
//			r := NewCompilationsHandler(test.useCaseMock, test.profileUsecaseMock, logger)
//			r.Register(server)
//			MCByMovieID := r.GetMCByMovieID()
//
//			_ = MCByMovieID(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result proto.MovieCompilation
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				for i, movie := range test.expected.Movies {
//					assert.Equal(t, movie.ID, result.Movies[i].ID)
//					assert.Equal(t, movie.Name, result.Movies[i].Name)
//					assert.Equal(t, movie.Picture, result.Movies[i].Picture)
//					for j, genre := range movie.Genre {
//						assert.Equal(t, genre.ID, result.Movies[i].Genre[j].ID)
//						assert.Equal(t, genre.Name, result.Movies[i].Genre[j].Name)
//					}
//				}
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
//
//func TestMoviesCompilationsDelivery_GetMCByGenreID(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movie1 := proto.MovieInfo{
//		ID:   1,
//		Name: "Movie1",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Боевик",
//			},
//			{
//				ID:   2,
//				Name: "Криминал",
//			},
//		},
//		Picture: "picture_name.webp",
//	}
//	movie2 := proto.MovieInfo{
//		ID:   2,
//		Name: "Movie2",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Фантастика",
//			},
//			{
//				ID:   2,
//				Name: "Семейный",
//			},
//		},
//		Picture: "picture_name2.webp",
//	}
//
//	MC1Movie1 := movie1
//	MC1Movie2 := movie2
//	MC1 := proto.MovieCompilation{
//		Name: "Test MC1",
//		Movies: []*proto.MovieInfo{
//			&MC1Movie1,
//			&MC1Movie2,
//		},
//	}
//
//	tests := []struct {
//		name               string
//		paramExists        bool
//		param              string
//		useCaseMock        *mock.MockMovieCompilationsClient
//		profileUsecaseMock *profile.MockProfileClient
//		expected           proto.MovieCompilation
//		expectedError      bool
//	}{
//		{
//			name:        "Get MC by genre ID",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByGenreFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expected:      MC1,
//			expectedError: false,
//		},
//		{
//			name:        "Return error",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByGenreFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return nil, errors.New(testError)
//				},
//			},
//			expectedError: true,
//		},
//		{
//			name:        "No param",
//			paramExists: false,
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByGenreFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expectedError: true,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			server := echo.New()
//			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/genre/1", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//			if test.paramExists {
//				ctx.SetParamNames("genre_id")
//				ctx.SetParamValues(test.param)
//			}
//
//			r := NewCompilationsHandler(test.useCaseMock, test.profileUsecaseMock, logger)
//			r.Register(server)
//			MCByMovieID := r.GetMCByGenre()
//
//			_ = MCByMovieID(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result proto.MovieCompilation
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				for i, movie := range test.expected.Movies {
//					assert.Equal(t, movie.ID, result.Movies[i].ID)
//					assert.Equal(t, movie.Name, result.Movies[i].Name)
//					assert.Equal(t, movie.Picture, result.Movies[i].Picture)
//					for j, genre := range movie.Genre {
//						assert.Equal(t, genre.ID, result.Movies[i].Genre[j].ID)
//						assert.Equal(t, genre.Name, result.Movies[i].Genre[j].Name)
//					}
//				}
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
//
//func TestMoviesCompilationsDelivery_GetMCByPersonID(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movie1 := proto.MovieInfo{
//		ID:   1,
//		Name: "Movie1",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Боевик",
//			},
//			{
//				ID:   2,
//				Name: "Криминал",
//			},
//		},
//		Picture: "picture_name.webp",
//	}
//	movie2 := proto.MovieInfo{
//		ID:   2,
//		Name: "Movie2",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Фантастика",
//			},
//			{
//				ID:   2,
//				Name: "Семейный",
//			},
//		},
//		Picture: "picture_name2.webp",
//	}
//
//	MC1Movie1 := movie1
//	MC1Movie2 := movie2
//	MC1 := proto.MovieCompilation{
//		Name: "Test MC1",
//		Movies: []*proto.MovieInfo{
//			&MC1Movie1,
//			&MC1Movie2,
//		},
//	}
//
//	tests := []struct {
//		name               string
//		paramExists        bool
//		param              string
//		useCaseMock        *mock.MockMovieCompilationsClient
//		profileUsecaseMock *profile.MockProfileClient
//		expected           proto.MovieCompilation
//		expectedError      bool
//	}{
//		{
//			name:        "Get MC by country ID",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByPersonFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expected:      MC1,
//			expectedError: false,
//		},
//		{
//			name:        "Return error",
//			paramExists: true,
//			param:       "1",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByPersonFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return nil, errors.New(testError)
//				},
//			},
//			expectedError: true,
//		},
//		{
//			name:        "No param",
//			paramExists: false,
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetByPersonFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expectedError: true,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			server := echo.New()
//			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/person/1", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//			if test.paramExists {
//				ctx.SetParamNames("person_id")
//				ctx.SetParamValues(test.param)
//			}
//
//			r := NewCompilationsHandler(test.useCaseMock, test.profileUsecaseMock, logger)
//			r.Register(server)
//			MCByMovieID := r.GetMCByPersonID()
//
//			_ = MCByMovieID(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result proto.MovieCompilation
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				for i, movie := range test.expected.Movies {
//					assert.Equal(t, movie.ID, result.Movies[i].ID)
//					assert.Equal(t, movie.Name, result.Movies[i].Name)
//					assert.Equal(t, movie.Picture, result.Movies[i].Picture)
//					for j, genre := range movie.Genre {
//						assert.Equal(t, genre.ID, result.Movies[i].Genre[j].ID)
//						assert.Equal(t, genre.Name, result.Movies[i].Genre[j].Name)
//					}
//				}
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
//
//func TestMoviesCompilationsDelivery_GetTopMC(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movie1 := proto.MovieInfo{
//		ID:   1,
//		Name: "Movie1",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Боевик",
//			},
//			{
//				ID:   2,
//				Name: "Криминал",
//			},
//		},
//		Picture: "picture_name.webp",
//	}
//	movie2 := proto.MovieInfo{
//		ID:   2,
//		Name: "Movie2",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Фантастика",
//			},
//			{
//				ID:   2,
//				Name: "Семейный",
//			},
//		},
//		Picture: "picture_name2.webp",
//	}
//
//	MC1Movie1 := movie1
//	MC1Movie2 := movie2
//	MC1 := proto.MovieCompilation{
//		Name: "Test MC1",
//		Movies: []*proto.MovieInfo{
//			&MC1Movie1,
//			&MC1Movie2,
//		},
//	}
//
//	tests := []struct {
//		name               string
//		paramExists        bool
//		param              string
//		useCaseMock        *mock.MockMovieCompilationsClient
//		profileUsecaseMock *profile.MockProfileClient
//		expected           proto.MovieCompilation
//		expectedError      bool
//	}{
//		{
//			name:        "Get top MC",
//			paramExists: true,
//			param:       "15",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetTopFunc: func(ctx context.Context, in *proto.GetCompilationOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expected:      MC1,
//			expectedError: false,
//		},
//		{
//			name:        "Return error",
//			paramExists: true,
//			param:       "10",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetTopFunc: func(ctx context.Context, in *proto.GetCompilationOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
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
//			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/top", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//
//			if test.paramExists {
//				ctx.QueryParams().Set("limit", test.param)
//			}
//
//			r := NewCompilationsHandler(test.useCaseMock, test.profileUsecaseMock, logger)
//			r.Register(server)
//			mainMC := r.GetTopMC()
//
//			_ = mainMC(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result proto.MovieCompilation
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				for i, movie := range test.expected.Movies {
//					assert.Equal(t, movie.ID, result.Movies[i].ID)
//					assert.Equal(t, movie.Name, result.Movies[i].Name)
//					assert.Equal(t, movie.Picture, result.Movies[i].Picture)
//					for j, genre := range movie.Genre {
//						assert.Equal(t, genre.ID, result.Movies[i].Genre[j].ID)
//						assert.Equal(t, genre.Name, result.Movies[i].Genre[j].Name)
//					}
//				}
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
//
//func TestMoviesCompilationsDelivery_GetTopByYear(t *testing.T) {
//	config := zap.NewDevelopmentConfig()
//	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	prLogger, _ := config.Build()
//	logger := prLogger.Sugar()
//	defer prLogger.Sync()
//	const testError = "test error"
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	movie1 := proto.MovieInfo{
//		ID:   1,
//		Name: "Movie1",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Боевик",
//			},
//			{
//				ID:   2,
//				Name: "Криминал",
//			},
//		},
//		Picture: "picture_name.webp",
//	}
//	movie2 := proto.MovieInfo{
//		ID:   2,
//		Name: "Movie2",
//		Genre: []*proto.Genre{
//			{
//				ID:   1,
//				Name: "Фантастика",
//			},
//			{
//				ID:   2,
//				Name: "Семейный",
//			},
//		},
//		Picture: "picture_name2.webp",
//	}
//
//	MC1Movie1 := movie1
//	MC1Movie2 := movie2
//	MC1 := proto.MovieCompilation{
//		Name: "Test MC1",
//		Movies: []*proto.MovieInfo{
//			&MC1Movie1,
//			&MC1Movie2,
//		},
//	}
//
//	tests := []struct {
//		name               string
//		paramExists        bool
//		param              string
//		useCaseMock        *mock.MockMovieCompilationsClient
//		profileUsecaseMock *profile.MockProfileClient
//		expected           proto.MovieCompilation
//		expectedError      bool
//	}{
//		{
//			name:        "Get Top by year",
//			paramExists: true,
//			param:       "2011",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetTopByYearFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expected:      MC1,
//			expectedError: false,
//		},
//		{
//			name:        "Return error",
//			paramExists: true,
//			param:       "2011",
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetTopByYearFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return nil, errors.New(testError)
//				},
//			},
//			expectedError: true,
//		},
//		{
//			name:        "No param",
//			paramExists: false,
//			useCaseMock: &mock.MockMovieCompilationsClient{
//				GetTopByYearFunc: func(ctx context.Context, in *proto.GetByIDOptions, opts ...grpc.CallOption) (*proto.MovieCompilation, error) {
//					return &MC1, nil
//				},
//			},
//			expectedError: true,
//		},
//	}
//
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			server := echo.New()
//			req := httptest.NewRequest(echo.GET, "/api/v1/movieCompilations/yearTop/2011", nil)
//			rec := httptest.NewRecorder()
//			ctx := server.NewContext(req, rec)
//			ctx.Set("REQUEST_ID", "1")
//			if test.paramExists {
//				ctx.SetParamNames("year")
//				ctx.SetParamValues(test.param)
//			}
//
//			r := NewCompilationsHandler(test.useCaseMock, test.profileUsecaseMock, logger)
//			r.Register(server)
//			MCByYear := r.GetYearTopMC()
//
//			_ = MCByYear(ctx)
//			body := rec.Body
//			status := rec.Result().Status
//			var result proto.MovieCompilation
//			_ = json.Unmarshal(body.Bytes(), &result)
//			if test.expectedError {
//				assert.Equal(t, "500 Internal Server Error", status)
//			} else {
//				for i, movie := range test.expected.Movies {
//					assert.Equal(t, movie.ID, result.Movies[i].ID)
//					assert.Equal(t, movie.Name, result.Movies[i].Name)
//					assert.Equal(t, movie.Picture, result.Movies[i].Picture)
//					for j, genre := range movie.Genre {
//						assert.Equal(t, genre.ID, result.Movies[i].Genre[j].ID)
//						assert.Equal(t, genre.Name, result.Movies[i].Genre[j].Name)
//					}
//				}
//				assert.Equal(t, "200 OK", status)
//			}
//		})
//	}
//}
