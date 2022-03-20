package main

import (

	"github.com/labstack/echo/v4"

	"log"
	"myapp/internal/composites"
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

	staffComposite, err := composites.NewStaffComposite(postgresDBC)
	if err != nil {
		log.Fatal("staff composite failed")
	}
	staffComposite.Handler.Register(echoServer)


	moviesCompilationsComposite, err := composites.NewMoviesCompilationsComposite(postgresDBC)
	if err != nil {
		log.Fatal("moviesCompilations composite failed")
	}
	moviesCompilationsComposite.Handler.Register(echoServer)

	userComposite, err := composites.NewUserComposite(postgresDBC, redisComposite)
	if err != nil {
		log.Fatal("user composite failed")
	}
	userComposite.Middleware.Register(echoServer)
	userComposite.Handler.Register(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))

}
