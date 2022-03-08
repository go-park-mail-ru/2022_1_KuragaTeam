package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"myapp/db"
	"myapp/handlers"
	"myapp/middleware"

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

	connRedis, err := db.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = connRedis.Close()
		if err != nil {
			log.Fatal(err)
		}
		dbPool.Close()
	}()

	e := echo.New()

	e.Use(middleware.CheckAuthorization(&connRedis))
	e.Use(middleware.CORS())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/api/v1/signup", handlers.CreateUserHandler(dbPool, &connRedis))
	e.POST("/api/v1/login", handlers.LoginUserHandler(dbPool, &connRedis))
	e.DELETE("/api/v1/logout", handlers.LogoutHandler(&connRedis))
	e.GET("/api/v1/", handlers.GetHomePageHandler(dbPool))

	e.Logger.Fatal(e.Start(":1323"))
}
