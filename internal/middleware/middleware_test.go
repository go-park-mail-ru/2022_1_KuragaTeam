package middleware

import (
	"errors"
	"myapp/internal/csrf"
	"myapp/internal/microservices/authorization/proto"
	"myapp/internal/microservices/authorization/usecase"
	"myapp/internal/monitoring/delivery"
	"net/http"
	"net/http/httptest"
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

func TestMiddleware_CheckAuthorization(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockAuthorizationClient(ctl)

	tests := []struct {
		name   string
		mock   func()
		cookie http.Cookie
		err    error
		userID int64
	}{
		{
			name: "User is Authorized",
			mock: func() {
				data := &proto.Cookie{Cookie: "session"}
				returnData := &proto.UserID{ID: int64(1)}
				gomock.InOrder(
					mockService.EXPECT().CheckAuthorization(gomock.Any(), data).Return(returnData, nil),
				)
			},
			cookie: http.Cookie{
				Name:       "Session_cookie",
				Value:      "session",
				Path:       "",
				Domain:     "",
				Expires:    time.Now().Add(time.Hour),
				RawExpires: "",
				MaxAge:     0,
				Secure:     false,
				HttpOnly:   true,
				SameSite:   0,
				Raw:        "",
				Unparsed:   nil,
			},
			err:    nil,
			userID: int64(1),
		},
		{
			name: "User is Unauthorized",
			mock: func() {
				data := &proto.Cookie{Cookie: "session"}
				returnData := &proto.UserID{ID: int64(-1)}
				gomock.InOrder(
					mockService.EXPECT().CheckAuthorization(gomock.Any(), data).Return(returnData, status.Error(codes.Internal, "error")),
				)
			},
			cookie: http.Cookie{
				Name:       "Session_cookie",
				Value:      "session",
				Path:       "",
				Domain:     "",
				Expires:    time.Now().Add(time.Hour),
				RawExpires: "",
				MaxAge:     0,
				Secure:     false,
				HttpOnly:   true,
				SameSite:   0,
				Raw:        "",
				Unparsed:   nil,
			},
			err:    errors.New("error"),
			userID: int64(-1),
		},
		{
			name: "No cookie in context",
			mock: func() {},
			cookie: http.Cookie{
				Name:     "Wrong_session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			userID: int64(-1),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(echo.GET, "/", nil)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Request().AddCookie(&th.cookie)

			th.mock()

			var metrics *delivery.PrometheusMetrics
			middleware := NewMiddleware(mockService, logger, metrics)
			middleware.Register(server)

			checkAuthorization := middleware.CheckAuthorization()
			handlerFunc := checkAuthorization(func(c echo.Context) error {
				return th.err
			})
			_ = handlerFunc(ctx)

			assert.Equal(t, ctx.Get("USER_ID"), th.userID)
		})
	}
}

func TestMiddleware_CSRF(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := usecase.NewMockAuthorizationClient(ctl)
	create, _ := csrf.Tokens.Create("session", time.Now().Add(time.Hour).Unix())

	tests := []struct {
		name           string
		cookie         http.Cookie
		expectedStatus int
		expectedJSON   string
		expectedError  bool
		method         string
		token          string
	}{
		{
			name: "No session cookie",
			cookie: http.Cookie{
				Name:     "Wrong_Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			method:         echo.PUT,
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
		},
		{
			name: "check csrf error",
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			method:         echo.PUT,
			expectedStatus: http.StatusInternalServerError,
			expectedError:  true,
			token:          "token",
		},
		{
			name: "not valid csrf",
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "wrong_session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			method:         echo.PUT,
			expectedStatus: http.StatusForbidden,
			expectedError:  true,
			token:          create,
		},
		{
			name: "valid csrf",
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			method:         echo.PUT,
			expectedError:  false,
			token:          create,
			expectedStatus: http.StatusOK,
		},
		{
			name: "method not put",
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			method:         echo.GET,
			expectedError:  false,
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := echo.New()
			th := test

			req := httptest.NewRequest(th.method, "/", nil)
			req.Header.Set("csrf-token", th.token)
			rec := httptest.NewRecorder()
			ctx := server.NewContext(req, rec)
			ctx.Request().AddCookie(&th.cookie)

			var metrics *delivery.PrometheusMetrics
			middleware := NewMiddleware(mockService, logger, metrics)
			middleware.Register(server)

			receivedCSRF := middleware.CSRF()
			handlerFunc := receivedCSRF(func(c echo.Context) error {
				return nil
			})
			err := handlerFunc(ctx)

			if th.expectedError == true {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, th.expectedStatus, rec.Code)
			}
		})
	}
}
