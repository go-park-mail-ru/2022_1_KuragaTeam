package middleware

import (
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"github.com/rafaeljusto/redigomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestCheckAuthorizationCase struct {
	name       string
	StatusCode int
	cookie     *http.Cookie
	userID     int64
}

func TestCheckAuthorization(t *testing.T) {
	cases := []TestCheckAuthorizationCase{
		TestCheckAuthorizationCase{
			name:       "Default logout",
			StatusCode: http.StatusOK,
			cookie: &http.Cookie{
				Name:     "Session_cookie",
				Value:    "okdg0wrijifoi",
				HttpOnly: true,
				Expires:  time.Now().Add(time.Hour),
				SameSite: 0,
				Secure:   false,
			},
			userID: 5,
		},
	}
	server := echo.New()
	defer func(server *echo.Echo) {
		err := server.Close()
		if err != nil {

		}
	}(server)

	conn := redigomock.NewConn()
	pool := &redis.Pool{
		// Return the same connection mock for each Get() call.
		Dial:        func() (redis.Conn, error) { return conn, nil },
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
	}

	for _, item := range cases {

		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		ctx := server.NewContext(req, rec)
		ctx.Request().AddCookie(item.cookie)

		cmd := conn.Command("GET", item.cookie.Value).Expect(item.userID)

		respMiddleware := CheckAuthorization(pool)
		resp := respMiddleware(func(c echo.Context) error {
			return nil
		})

		if assert.NoError(t, resp(ctx)) {
			assert.Equal(t, conn.Stats(cmd), 1)
		}
		assert.Equal(t, item.StatusCode, rec.Code)

		assert.Equal(t, ctx.Get("USER_ID"), item.userID)
	}
}
