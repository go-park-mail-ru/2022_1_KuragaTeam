package composites

import (
	"go.uber.org/zap"
	countryRepository "myapp/internal/country/repository"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/microservices/movie"
	movieRepository "myapp/internal/microservices/movie/repository"
	"myapp/internal/microservices/movie/usecase"
	personsRepository "myapp/internal/persons/repository"
)

type MovieComposite struct {
	Storage movie.Storage
	Service *usecase.Service
}

func NewMovieComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*MovieComposite, error) {
	movieStorage := movieRepository.NewStorage(postgresComposite.Db)
	genreStorage := genreRepository.NewStorage(postgresComposite.Db)
	countryStorage := countryRepository.NewStorage(postgresComposite.Db)
	staffStorage := personsRepository.NewStorage(postgresComposite.Db)
	service := usecase.NewService(movieStorage, genreStorage, countryStorage, staffStorage)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
	}, nil
}
