package middleware

import (
	"myapp/internal/csrf"
	"myapp/internal/user"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Middleware struct {
	userService user.Service
	logger      *zap.SugaredLogger
}

func NewMiddleware(service user.Service, logger *zap.SugaredLogger) *Middleware {
	return &Middleware{
		userService: service,
		logger:      logger,
	}
}

func (m Middleware) Register(router *echo.Echo) {
	router.Use(m.CheckAuthorization())
	router.Use(m.CORS())
	router.Use(m.AccessLog())
	router.Use(m.CSRF())
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
					ctx.Set("USER_ID", int64(-1))
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
		AllowOrigins:     []string{"http://movie-space.ru:8080", "http://localhost:8080", "http://localhost:1323"},
		AllowHeaders:     []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With", "csrf-token"},
		AllowCredentials: true,
		MaxAge:           84600,
	})
}

func (m Middleware) AccessLog() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			id, _ := uuid.NewV4()

			start := time.Now()
			ctx.Set("REQUEST_ID", id.String())

			m.logger.Info(
				zap.String("ID", id.String()),
				zap.String("URL", ctx.Request().URL.Path),
				zap.String("METHOD", ctx.Request().Method),
			)

			err := next(ctx)

			responseTime := time.Since(start)
			m.logger.Info(
				zap.String("ID", id.String()),
				zap.Duration("TIME FOR ANSWER", responseTime),
			)

			return err
		}
	}
}

func (m Middleware) CSRF() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if ctx.Request().Method == "PUT" || ctx.Request().RequestURI == "/api/v1/edit" {
				cookie, err := ctx.Cookie("Session_cookie")
				if err != nil {
					m.logger.Debug(
						zap.String("COOKIE", err.Error()),
						zap.Int("ANSWER STATUS", http.StatusInternalServerError),
					)

					return ctx.JSON(http.StatusInternalServerError, &user.Response{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}

				GetToken := ctx.Request().Header.Get("csrf-token")

				isValidCsrf, err := csrf.Tokens.Check(cookie.Value, GetToken)
				if err != nil {
					return ctx.JSON(http.StatusInternalServerError, &user.Response{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}

				if !isValidCsrf {
					return ctx.JSON(http.StatusForbidden, &user.Response{
						Status:  http.StatusForbidden,
						Message: "Wrong csrf token",
					})
				}
			}
			return next(ctx)
		}
	}
}
