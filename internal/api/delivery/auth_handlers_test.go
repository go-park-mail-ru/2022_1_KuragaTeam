package delivery

import (
	"myapp/internal/constants"
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/microservices/authorization/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestAuthDelivery_SignUp(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockAuthorizationClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		data           string
		requestID      string
	}{
		{
			name: "Handler returned status 201",
			mock: func() {
				userData := &proto.SignUpData{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				returnData := &proto.Cookie{Cookie: "session"}
				gomock.InOrder(
					mockService.EXPECT().SignUp(gomock.Any(), userData).Return(returnData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"User created\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 400, Email is not unique",
			mock: func() {
				userData := &proto.SignUpData{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().SignUp(gomock.Any(), userData).Return(nil, status.Error(codes.InvalidArgument, constants.ErrEmailIsNotUnique.Error())),
				)
			},
			expectedStatus: http.StatusBadRequest,
			expectedJSON:   "{\"status\":400,\"message\":\"" + constants.ErrEmailIsNotUnique.Error() + "\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase SignUp error",
			mock: func() {
				userData := &proto.SignUpData{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().SignUp(gomock.Any(), userData).Return(nil, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, wrong REQUEST_ID",
			mock: func() {
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "WRONG_REQUEST_ID",
		},
		{
			name: "Handler returned status 500, bad connection",
			mock: func() {
				userData := &proto.SignUpData{
					Name:     "Olga",
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().SignUp(gomock.Any(), userData).Return(nil, status.Error(codes.Unavailable, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.POST, "/api/v1/signup", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")

			th.mock()

			handler := NewAuthHandler(logger, mockService)
			handler.Register(server)

			signup := handler.SignUp()
			_ = signup(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestAuthHandler_LogIn(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockAuthorizationClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		data           string
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.LogInData{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				returnData := &proto.Cookie{Cookie: "session"}
				gomock.InOrder(
					mockService.EXPECT().LogIn(gomock.Any(), userData).Return(returnData, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"User can be logged in\"}",
			data:           `{"email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 404, User not found",
			mock: func() {
				userData := &proto.LogInData{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().LogIn(gomock.Any(), userData).Return(nil, status.Error(codes.NotFound, constants.ErrWrongData.Error())),
				)
			},
			expectedStatus: http.StatusNotFound,
			expectedJSON:   "{\"status\":404,\"message\":\"" + constants.ErrWrongData.Error() + "\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase LogIn error",
			mock: func() {
				userData := &proto.LogInData{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().LogIn(gomock.Any(), userData).Return(nil, status.Error(codes.Internal, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"username": "Olga", "email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, wrong REQUEST_ID",
			mock: func() {
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			requestID:      "WRONG_REQUEST_ID",
			data:           `{"email": "olga@mail.ru", "password": "olga123321"}`,
		},
		{
			name: "Handler returned status 500, bad connection",
			mock: func() {
				userData := &proto.LogInData{
					Email:    "olga@mail.ru",
					Password: "olga123321",
				}
				gomock.InOrder(
					mockService.EXPECT().LogIn(gomock.Any(), userData).Return(nil, status.Error(codes.Unavailable, "error")),
				)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			data:           `{"email": "olga@mail.ru", "password": "olga123321"}`,
			requestID:      "REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.POST, "/api/v1/login", strings.NewReader(th.data))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")

			th.mock()

			handler := NewAuthHandler(logger, mockService)
			handler.Register(server)

			login := handler.LogIn()
			_ = login(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}

func TestAuthHandler_LogOut(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockAuthorizationClient(ctl)

	tests := []struct {
		name           string
		mock           func()
		expectedStatus int
		expectedJSON   string
		cookie         http.Cookie
		requestID      string
	}{
		{
			name: "Handler returned status 200",
			mock: func() {
				userData := &proto.Cookie{
					Cookie: "session",
				}
				gomock.InOrder(
					mockService.EXPECT().LogOut(gomock.Any(), userData).Return(&proto.Empty{}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedJSON:   "{\"status\":200,\"message\":\"User is logged out\"}",
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			requestID: "REQUEST_ID",
		},
		{
			name:           "Handler returned status 500, wrong cookie",
			mock:           func() {},
			expectedStatus: http.StatusInternalServerError,
			cookie: http.Cookie{
				Name:     "Wrong_Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			expectedJSON: "{\"status\":500,\"message\":\"http: named cookie not present\"}",
			requestID:    "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, usecase LogOut error",
			mock: func() {
				userData := &proto.Cookie{
					Cookie: "session",
				}
				gomock.InOrder(
					mockService.EXPECT().LogOut(gomock.Any(), userData).Return(&proto.Empty{}, status.Error(codes.Internal, "error")),
				)
			},
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"error\"}",
			requestID:      "REQUEST_ID",
		},
		{
			name: "Handler returned status 500, wrong REQUEST_ID",
			mock: func() {
			},
			expectedStatus: http.StatusInternalServerError,
			expectedJSON:   "{\"status\":500,\"message\":\"" + constants.NoRequestID + "\"}",
			requestID:      "WRONG_REQUEST_ID",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.DELETE, "/api/v1/logout", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Set(th.requestID, "1")

			ctx.Request().AddCookie(&th.cookie)

			th.mock()

			handler := NewAuthHandler(logger, mockService)
			handler.Register(server)

			logout := handler.LogOut()
			_ = logout(ctx)

			body := rec.Body.String()
			status := rec.Code

			assert.Equal(t, test.expectedJSON, body)
			assert.Equal(t, th.expectedStatus, status)
		})
	}
}
