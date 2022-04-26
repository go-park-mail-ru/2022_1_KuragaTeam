package main

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"myapp/internal/api/delivery"
	"myapp/internal/composites"
	compilationsPB "myapp/internal/microservices/compilations/proto"
	pb "myapp/internal/microservices/movie/proto"
	"os"
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
		logger.Fatal("minio composite failed", err)
	}

	conn, err := grpc.Dial(
		os.Getenv("MOVIE_HOST")+":"+os.Getenv("MOVIE_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMoviesClient(conn)

	movieHandler := delivery.NewMovieHandler(c, logger)

	movieHandler.Register(echoServer)

	staffComposite, err := composites.NewStaffComposite(postgresDBC, logger)
	if err != nil {
		logger.Fatal("staff composite failed")
	}
	staffComposite.Handler.Register(echoServer)

	conn2, err := grpc.Dial(
		os.Getenv("COMPILATIONS_HOST")+":"+os.Getenv("COMPILATIONS_PORT"),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c2 := compilationsPB.NewMovieCompilationsClient(conn2)

	compilationsHandler := delivery.NewHandler(c2, logger)

	compilationsHandler.Register(echoServer)

	//moviesCompilationsComposite, err := composites.NewMoviesCompilationsComposite(postgresDBC, logger)
	//if err != nil {
	//	logger.Fatal("moviesCompilations composite failed")
	//}
	//moviesCompilationsComposite.Service.Register(echoServer)

	userComposite, err := composites.NewUserComposite(postgresDBC, redisComposite, minioComposite, logger)
	if err != nil {
		logger.Fatal("user composite failed")
	}

	userComposite.Middleware.Register(echoServer)
	userComposite.Handler.Register(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))
}
