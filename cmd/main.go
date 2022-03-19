package main

import (
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
  
	"log"
	"myapp/db"
	"myapp/handlers"
	"myapp/internal/composites"
	"myapp/middleware"
	"myapp/utils"
)

// @title Movie Space API
// @version 1.0
// @description This is API server for Movie Space website.
// @license.name  ""

// @host movie-space.ru:1323
// @BasePath /api/v1/
// @schemes http
func main() {
	echoServer := echo.New()

	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		log.Fatal("postgresdb composite failed")
	}

	redisComposite, err := composites.NewRedisComposite()
	if err != nil {
		log.Fatal("redis composite failed")
	}

	movieComposite, err := composites.NewMovieComposite(postgresDBC)
	if err != nil {
		log.Fatal("author composite failed")
	}

  movieComposite.Handler.Register(echoServer)

	userComposite, err := composites.NewUserComposite(postgresDBC, redisComposite)
	if err != nil {
		log.Fatal("user composite failed")
	}
	userComposite.Middleware.Register(echoServer)
	userComposite.Handler.Register(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))

}
