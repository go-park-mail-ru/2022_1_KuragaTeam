package main

import (
	"log"
	"myapp/internal/composites"
	"myapp/internal/microservices/movie/proto"
	"net"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	prLogger, err := config.Build()
	if err != nil {
		log.Fatal("zap logger build error")
	}

	logger := prLogger.Sugar()
	defer func(prLogger *zap.Logger) {
		err = prLogger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(prLogger)

	postgresDBC, err := composites.NewPostgresDBComposite()
	if err != nil {
		logger.Fatal("postgres db composite failed")
	}

	lis, err := net.Listen("tcp", ":"+os.Getenv("MOVIE_PORT"))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	grpcServ := grpc.NewServer()

	composite, err := composites.NewMovieComposite(postgresDBC, logger)
	if err != nil {
		return
	}

	proto.RegisterMoviesServer(grpcServ, composite.Service)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServ.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
