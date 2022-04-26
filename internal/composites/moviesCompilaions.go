package composites

import (
	"go.uber.org/zap"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/microservices/compilations"
	"myapp/internal/microservices/compilations/repository"
	"myapp/internal/microservices/compilations/usecase"
)

type MoviesCompilationsComposite struct {
	Service *usecase.Service
	Storage compilations.Storage
}

func NewMoviesCompilationsComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*MoviesCompilationsComposite, error) {
	MCStorage := repository.NewStorage(postgresComposite.Db)
	genreStorage := genreRepository.NewStorage(postgresComposite.Db)
	service := usecase.NewService(MCStorage, genreStorage)
	return &MoviesCompilationsComposite{
		Service: service,
		Storage: MCStorage,
	}, nil
}
