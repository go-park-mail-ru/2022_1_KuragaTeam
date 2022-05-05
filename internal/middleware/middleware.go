package middleware

import (
	"context"
	"myapp/internal/csrf"
	authorization "myapp/internal/microservices/authorization/proto"
	"myapp/internal/models"
	"myapp/internal/monitoring/delivery"
	"net/http"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Middleware struct {
	authMicroservice authorization.AuthorizationClient
	logger           *zap.SugaredLogger
	metrics          *delivery.PrometheusMetrics
}

func NewMiddleware(authMicroservice authorization.AuthorizationClient, logger *zap.SugaredLogger, monitoring *delivery.PrometheusMetrics) *Middleware {
	return &Middleware{
		authMicroservice: authMicroservice,
		logger:           logger,
		metrics:          monitoring,
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
				data := &authorization.Cookie{Cookie: cookie.Value}
				id, err := m.authMicroservice.CheckAuthorization(context.Background(), data)
				if err != nil {
					cookie = &http.Cookie{Expires: time.Now().AddDate(0, 0, -1)}
					ctx.SetCookie(cookie)
					ctx.Set("USER_ID", int64(-1))
					return next(ctx)
				}
				userID = id.ID
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
		AllowOrigins:     []string{"http://movie-space.ru:8080", "http://localhost:8080", "http://movie-space.ru", "https://movie-space.ru"},
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

			status := strconv.Itoa(ctx.Response().Status)
			path := ctx.Request().URL.Path
			method := ctx.Request().Method

			m.metrics.Hits.WithLabelValues(status, path, method).Inc()
			m.metrics.Duration.WithLabelValues(status, path, method).Observe(responseTime.Seconds())

			return err
		}
	}
}

func (m Middleware) CSRF() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if ctx.Request().Method == "PUT" {
				cookie, err := ctx.Cookie("Session_cookie")
				if err != nil {
					m.logger.Debug(
						zap.String("COOKIE", err.Error()),
						zap.Int("ANSWER STATUS", http.StatusInternalServerError),
					)

					return ctx.JSON(http.StatusInternalServerError, &models.Response{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}

				GetToken := ctx.Request().Header.Get("csrf-token")

				isValidCsrf, err := csrf.Tokens.Check(cookie.Value, GetToken)
				if err != nil {
					return ctx.JSON(http.StatusInternalServerError, &models.Response{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}

				if !isValidCsrf {
					return ctx.JSON(http.StatusForbidden, &models.Response{
						Status:  http.StatusForbidden,
						Message: "Wrong csrf token",
					})
				}
			}
			return next(ctx)
		}
	}
}
