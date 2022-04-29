package composites

import (
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/microservices/compilations"
	"myapp/internal/microservices/compilations/repository"
	"myapp/internal/microservices/compilations/usecase"
	personRepository "myapp/internal/persons/repository"

	"go.uber.org/zap"
)

type MoviesCompilationsComposite struct {
	Service *usecase.Service
	Storage compilations.Storage
}

func NewMoviesCompilationsComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*MoviesCompilationsComposite, error) {
	MCStorage := repository.NewStorage(postgresComposite.db)
	genreStorage := genreRepository.NewStorage(postgresComposite.db)
	personStorage := personRepository.NewStorage(postgresComposite.db)
	service := usecase.NewService(MCStorage, genreStorage, personStorage)
	return &MoviesCompilationsComposite{
		Service: service,
		Storage: MCStorage,
	}, nil
}
