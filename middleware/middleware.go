package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4/middleware"

	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo/v4"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://movie-space.ru:8080", "http://localhost:8080"},
		AllowHeaders:     []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
	})
}

func CheckAuthorization(redisPool *redis.Pool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("Session_cookie")
			var userID int64
			userID = -1
			if err == nil {
				connRedis := redisPool.Get()
				defer connRedis.Close()
				userID, err = redis.Int64(connRedis.Do("GET", cookie.Value))
				if err != nil {
					cookie = &http.Cookie{Expires: time.Now().AddDate(0, 0, -1)}
					ctx.SetCookie(cookie)
					ctx.Set("USER_ID", -1)
					log.Println(err)
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
