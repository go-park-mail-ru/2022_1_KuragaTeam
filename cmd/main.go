package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"myapp/internal/api/delivery"
	"myapp/internal/composites"
	pb "myapp/internal/microservices/movie/proto"
)

var (
	addr = flag.String("addr", "movie:5001", "the address to connect to")
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

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMoviesClient(conn)

	movieHandler := delivery.NewMovieHandler(c, logger)

	movieHandler.Register(echoServer)

	//movieComposite, err := composites.NewMovieComposite(postgresDBC, logger)
	//if err != nil {
	//	logger.Fatal("author composite failed")
	//}
	//movieComposite.Handler.Register(echoServer)

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
