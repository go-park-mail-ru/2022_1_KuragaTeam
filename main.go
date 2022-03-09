package main

import (
	"log"
	"myapp/db"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/jackc/pgx/v4"
)

// @title Movie Space API
// @version 1.0
// @description This is API server for Movie Space website.
// @license.name  ""

// @host movie-space.ru:1323
// @BasePath /api/v1/
// @schemes http
func main() {
	dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	redisPool, err := db.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()

	e := echo.New()

	e.Use(middleware.CheckAuthorization(redisPool))
	e.Use(middleware.CORS())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/api/v1/signup", handlers.CreateUserHandler(dbPool, redisPool))
	e.POST("/api/v1/login", handlers.LoginUserHandler(dbPool, redisPool))
	e.DELETE("/api/v1/logout", handlers.LogoutHandler(redisPool))
	e.GET("/api/v1/", handlers.GetHomePageHandler(dbPool))
	e.GET("/api/v1/movieCompilations", handlers.GetMovieCompilations(dbPool))

	e.Logger.Fatal(e.Start(":1323"))
}
