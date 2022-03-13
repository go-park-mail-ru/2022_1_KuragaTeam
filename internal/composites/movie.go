package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/api/movie"
	movie3 "myapp/internal/adapters/db/movie"
	movie2 "myapp/internal/domain/movie"
)

type MovieComposite struct {
	Storage movie2.Storage
	Service movie.Service
	Handler api.Handler
}

func NewMovieComposite(postgresComposite *PostgresDBComposite) (*MovieComposite, error) {
	storage := movie3.NewStorage(postgresComposite.db)
	service := movie2.NewService(storage)
	handler := movie.NewHandler(service)
	return &MovieComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
