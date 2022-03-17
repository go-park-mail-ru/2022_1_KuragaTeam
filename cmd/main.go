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
	//dbPool, err := db.ConnectDB()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer dbPool.Close()

	redisPool, err := db.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}
	defer func(redisPool *redis.Pool) {
		err = redisPool.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(redisPool)

	echoServ := echo.New()
	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		log.Fatal("postgresdb composite failed")
	}

	userPool := utils.UserPool{
		Pool: (*postgresDBC).DB,
	}
	movieComposite, err := composites.NewMovieComposite(postgresDBC)
	if err != nil {
		log.Fatal("author composite failed")
	}
	movieComposite.Handler.Register(echoServ)

	moviesCompilationsComposite, err := composites.NewMoviesCompilationsComposite(postgresDBC)
	moviesCompilationsComposite.Handler.Register(echoServ)

	echoServ.Use(middleware.CheckAuthorization(redisPool))
	echoServ.Use(middleware.CORS())
	echoServ.GET("/swagger/*", echoSwagger.WrapHandler)

	echoServ.POST("/api/v1/signup", handlers.CreateUserHandler(&userPool, redisPool))
	echoServ.POST("/api/v1/login", handlers.LoginUserHandler(&userPool, redisPool))
	echoServ.DELETE("/api/v1/logout", handlers.LogoutHandler(redisPool))
	echoServ.GET("/api/v1/", handlers.GetHomePageHandler(&userPool))

	echoServ.GET("/api/v1/movieCompilations", handlers.GetMovieCompilations())

	echoServ.Logger.Fatal(echoServ.Start(":1323"))
}
