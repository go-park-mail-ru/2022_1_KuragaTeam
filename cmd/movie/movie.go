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
	lis, err := net.Listen("tcp", ":"+os.Getenv("MOVIE_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

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

	composite, err := composites.NewMovieComposite(postgresDBC, logger)
	if err != nil {
		return
	}

	proto.RegisterMoviesServer(s, composite.Service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
