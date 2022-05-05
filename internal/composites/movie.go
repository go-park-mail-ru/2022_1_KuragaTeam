package composites

import (
	countryRepository "myapp/internal/country/repository"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/microservices/movie"
	movieRepository "myapp/internal/microservices/movie/repository"
	"myapp/internal/microservices/movie/usecase"
	personsRepository "myapp/internal/persons/repository"

	"go.uber.org/zap"
)

type MovieComposite struct {
	Storage movie.Storage
	Service *usecase.Service
}

func NewMovieComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*MovieComposite, error) {
	movieStorage := movieRepository.NewStorage(postgresComposite.db)
	genreStorage := genreRepository.NewStorage(postgresComposite.db)
	countryStorage := countryRepository.NewStorage(postgresComposite.db)
	staffStorage := personsRepository.NewStorage(postgresComposite.db)
	service := usecase.NewService(movieStorage, genreStorage, countryStorage, staffStorage)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
	}, nil
}
