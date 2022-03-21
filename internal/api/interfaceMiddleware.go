package api

import "github.com/labstack/echo/v4"

type Middleware interface {
	Register(router *echo.Echo)
}
