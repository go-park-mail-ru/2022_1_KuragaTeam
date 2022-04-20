package main

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"myapp/internal/composites"
	countryRepository "myapp/internal/country/repository"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/microservices/movie/proto"
	"myapp/internal/microservices/movie/repository"
	"myapp/internal/microservices/movie/usecase"
	personsRepository "myapp/internal/persons/repository"

	"net"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", ":5001")
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

	movieStorage := repository.NewStorage(postgresDBC.Db)
	genreStorage := genreRepository.NewStorage(postgresDBC.Db)
	countryStorage := countryRepository.NewStorage(postgresDBC.Db)
	staffStorage := personsRepository.NewStorage(postgresDBC.Db)

	proto.RegisterMoviesServer(s, usecase.NewService(movieStorage, genreStorage, countryStorage, staffStorage))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
