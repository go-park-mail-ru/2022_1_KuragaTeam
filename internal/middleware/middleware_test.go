package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"myapp/internal/user/delivery/mock"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestMiddleware_CheckAuthorization(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, _ := config.Build()
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	mockService := mock.NewMockService(ctl)

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
				gomock.InOrder(
					mockService.EXPECT().CheckAuthorization("session").Return(int64(1), nil),
				)
			},
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
			},
			err:    nil,
			userID: int64(1),
		},
		{
			name: "User is Unauthorized",
			mock: func() {
				gomock.InOrder(
					mockService.EXPECT().CheckAuthorization("session").Return(int64(-1), errors.New("error")),
				)
			},
			cookie: http.Cookie{
				Name:     "Session_cookie",
				Value:    "session",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
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

			middleware := NewMiddleware(mockService, logger)
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