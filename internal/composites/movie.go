package composites

import (
	"go.uber.org/zap"
	"myapp/internal/api"
	countryRepository "myapp/internal/country/repository"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/movie"
	"myapp/internal/movie/delivery"
	"myapp/internal/movie/repository"
	"myapp/internal/movie/usecase"
	personsRepository "myapp/internal/persons/repository"
)

type MovieComposite struct {
	Storage movie.Storage
	Service movie.Service
	Handler api.Handler
}

func NewMovieComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*MovieComposite, error) {
	movieStorage := repository.NewStorage(postgresComposite.db)
	genreStorage := genreRepository.NewStorage(postgresComposite.db)
	countryStorage := countryRepository.NewStorage(postgresComposite.db)
	staffStorage := personsRepository.NewStorage(postgresComposite.db)
	service := usecase.NewService(movieStorage, genreStorage, countryStorage, staffStorage)
	handler := delivery.NewHandler(service, logger)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
		Handler: handler,
	}, nil
}
