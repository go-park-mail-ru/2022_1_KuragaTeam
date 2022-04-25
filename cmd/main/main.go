package main

import (
	"log"
	api "myapp/internal/api/delivery"
	"myapp/internal/composites"
	authMicroservice "myapp/internal/microservices/authorization/proto"
	profileMicroservice "myapp/internal/microservices/profile/proto"
	"myapp/internal/middleware"
	"os"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func LoadMicroservices(server *echo.Echo) (authMicroservice.AuthorizationClient,
	profileMicroservice.ProfileClient, []*grpc.ClientConn) {
	connections := make([]*grpc.ClientConn, 0)

	authConn, err := grpc.Dial(
		os.Getenv("AUTH_HOST")+":"+
			os.Getenv("AUTH_PORT"),
		grpc.WithInsecure(),
	)

	if err != nil {
		server.Logger.Fatal("authorization cant connect to grpc")
	}
	connections = append(connections, authConn)

	authorizationManager := authMicroservice.NewAuthorizationClient(authConn)

	profileConn, err := grpc.Dial(
		os.Getenv("PROFILE_HOST")+":"+
			os.Getenv("PROFILE_PORT"),
		grpc.WithInsecure(),
	)

	if err != nil {
		server.Logger.Fatal("profile cant connect to grpc")
	}
	connections = append(connections, profileConn)

	profileManager := profileMicroservice.NewProfileClient(profileConn)

	return authorizationManager, profileManager, connections
}

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

	auth, profile, conn := LoadMicroservices(echoServer)
	defer func() {
		if len(conn) == 0 {
			return
		}
		for _, c := range conn {
			err := c.Close()
			if err != nil {
				log.Fatalf("Error occurred during closing connection: %s", err.Error())
			}
		}
	}()

	appHandler := api.NewAPIMicroservices(logger, auth, profile)
	appHandler.Register(echoServer)

	///////////////////////////////////////////////////////

	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		logger.Fatal("postgres db composite failed")
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

	middlwares := middleware.NewMiddleware(auth, logger)
	middlwares.Register(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":1323"))

}
