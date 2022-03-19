package user

import (
	"myapp/internal/adapters/api"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
	userService Service
}

func NewMiddleware(service Service) api.Middleware {
	return &Middleware{userService: service}
}

func (m Middleware) Register(router *echo.Echo) {
	router.Use(m.CheckAuthorization())
	router.Use(m.CORS())
}

func (m Middleware) CheckAuthorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("Session_cookie")
			var userID int64
			userID = -1
			if err == nil {
				userID, err = m.userService.CheckAuthorization(cookie.Value)
				if err != nil {
					cookie = &http.Cookie{Expires: time.Now().AddDate(0, 0, -1)}
					ctx.SetCookie(cookie)
					ctx.Set("USER_ID", -1)
					return next(ctx)
				}
			}
			if err != nil {
				cookie = &http.Cookie{Expires: time.Now().AddDate(0, 0, -1)}
				ctx.SetCookie(cookie)
			}

			ctx.Set("USER_ID", userID)

			return next(ctx)
		}
	}
}

func (m Middleware) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://movie-space.ru:8080", "http://localhost:8080"},
		AllowHeaders:     []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           84600,
	})
}
