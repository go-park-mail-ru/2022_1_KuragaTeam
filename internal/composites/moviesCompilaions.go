package composites

import (
	"myapp/internal/api"
	"myapp/internal/moviesCompilations"
	"myapp/internal/moviesCompilations/delivery"
	"myapp/internal/moviesCompilations/usecase"
)

type MoviesCompilationsComposite struct {
	Service moviesCompilations.Service
	Handler api.Handler
}

func NewMoviesCompilationsComposite(postgresComposite *PostgresDBComposite) (*MoviesCompilationsComposite, error) {
	service := usecase.NewService()
	handler := delivery.NewHandler(service)
	return &MoviesCompilationsComposite{
		Service: service,
		Handler: handler,
	}, nil
}
