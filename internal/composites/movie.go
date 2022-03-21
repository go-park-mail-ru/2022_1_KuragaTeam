package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/db/country"
	"myapp/internal/adapters/db/genre"
	"myapp/internal/adapters/db/staff"
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
	genreStorage := genre.NewStorage(postgresComposite.db)
	countryStorage := country.NewStorage(postgresComposite.db)
	staffStorage := staff.NewStorage(postgresComposite.db)
	service := usecase.NewService(movieStorage, genreStorage, countryStorage, staffStorage)
	handler := delivery.NewHandler(service)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
		Handler: handler,
	}, nil
}
