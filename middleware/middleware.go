package middleware

import (
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/labstack/echo/v4"
)

func CheckAuthorization(connRedis *redis.Conn) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			cookie, err := ctx.Cookie("Session_cookie")
			var userID int64
			userID = -1
			if err == nil {
				userID, err = redis.Int64((*connRedis).Do("GET", cookie.Value))
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
