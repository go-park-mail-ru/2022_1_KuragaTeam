package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/api/movie"
	"myapp/internal/adapters/db/country"
	"myapp/internal/adapters/db/genre"
	movie3 "myapp/internal/adapters/db/movie"
	movie2 "myapp/internal/domain/movie"
)

type MovieComposite struct {
	Storage movie2.Storage
	Service movie.Service
	Handler api.Handler
}

func NewMovieComposite(postgresComposite *PostgresDBComposite) (*MovieComposite, error) {
	movieStorage := movie3.NewStorage(postgresComposite.db)
	genreStorage := genre.NewStorage(postgresComposite.db)
	countryStorage := country.NewStorage(postgresComposite.db)
	service := movie2.NewService(movieStorage, genreStorage, countryStorage)
	handler := movie.NewHandler(service)
	return &MovieComposite{
		Storage: movieStorage,
		Service: service,
		Handler: handler,
	}, nil
}
