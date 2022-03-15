package composites

import (
	"myapp/internal/adapters/api"
	moviesCompilations2 "myapp/internal/adapters/api/moviesCompilations"
	"myapp/internal/domain/moviesCompilations"
)

type MoviesCompilationsComposite struct {
	Service moviesCompilations2.Service
	Handler api.Handler
}

func NewMoviesCompilationsComposite(postgresComposite *PostgresDBComposite) (*MoviesCompilationsComposite, error) {
	service := moviesCompilations.NewService()
	handler := moviesCompilations2.NewHandler(service)
	return &MoviesCompilationsComposite{
		Service: service,
		Handler: handler,
	}, nil
}
