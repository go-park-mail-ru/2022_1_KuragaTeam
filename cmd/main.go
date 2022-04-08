package main

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, err := config.Build()
	if err != nil {
		log.Fatal("zap logger build error")
	}
	logger := prLogger.Sugar()
	defer prLogger.Sync()

	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		logger.Fatal("postgres db composite failed")
	}

	redisComposite, err := composites.NewRedisComposite()
	if err != nil {
		logger.Fatal("redis composite failed")
	}

	minioComposite, err := composites.NewMinioComposite()
	if err != nil {
		logger.Fatal("minio composite failed")
	}

	movieComposite, err := composites.NewMovieComposite(postgresDBC, logger)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	movieComposite.Handler.Register(echoServer)

	staffComposite, err := composites.NewStaffComposite(postgresDBC, logger)
	if err != nil {
		logger.Fatal("staff composite failed")
	}
	staffComposite.Handler.Register(echoServer)

	moviesCompilationsComposite, err := composites.NewMoviesCompilationsComposite(postgresDBC, logger)
	if err != nil {
		logger.Fatal("moviesCompilations composite failed")
	}
	moviesCompilationsComposite.Handler.Register(echoServer)

	userComposite, err := composites.NewUserComposite(postgresDBC, redisComposite, minioComposite, logger)
	if err != nil {
		logger.Fatal("user composite failed")
	}

	userComposite.Middleware.Register(echoServer)
	userComposite.Handler.Register(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))

}
