package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/api/movie"
	"myapp/internal/adapters/db/country"
	"myapp/internal/adapters/db/genre"
	movie3 "myapp/internal/adapters/db/movie"
	"myapp/internal/adapters/db/staff"
	movie2 "myapp/internal/domain/movie"
)

type MovieComposite struct {
	Storage movie2.Storage
	Service movie.Service
	Handler api.Handler
}

func NewMovieComposite(postgresComposite *PostgresDBComposite) (*MovieComposite, error) {
	movieStorage := movie3.NewStorage(postgresComposite.DB)
	genreStorage := genre.NewStorage(postgresComposite.DB)
	countryStorage := country.NewStorage(postgresComposite.DB)
	staffStorage := staff.NewStorage(postgresComposite.DB)
	service := movie2.NewService(movieStorage, genreStorage, countryStorage, staffStorage)
	handler := movie.NewHandler(service)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
		Handler: handler,
	}, nil
}
