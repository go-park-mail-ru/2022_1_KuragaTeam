package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/db/staff"
	countryRepository "myapp/internal/country/repository"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/movie"
	"myapp/internal/movie/delivery"
	"myapp/internal/movie/repository"
	"myapp/internal/movie/usecase"
)

type MovieComposite struct {
	Storage movie.Storage
	Service movie.Service
	Handler api.Handler
}

func NewMovieComposite(postgresComposite *PostgresDBComposite) (*MovieComposite, error) {
	movieStorage := repository.NewStorage(postgresComposite.db)
	genreStorage := genreRepository.NewStorage(postgresComposite.db)
	countryStorage := countryRepository.NewStorage(postgresComposite.db)
	staffStorage := staff.NewStorage(postgresComposite.db)
	service := usecase.NewService(movieStorage, genreStorage, countryStorage, staffStorage)
	handler := delivery.NewHandler(service)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
		Handler: handler,
	}, nil
}
