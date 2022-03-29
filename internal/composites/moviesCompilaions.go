package composites

import (
	"myapp/internal/api"
	genreRepository "myapp/internal/genre/repository"
	"myapp/internal/moviesCompilations"
	"myapp/internal/moviesCompilations/delivery"
	"myapp/internal/moviesCompilations/repository"
	"myapp/internal/moviesCompilations/usecase"
)

type MoviesCompilationsComposite struct {
	Service moviesCompilations.Service
	Handler api.Handler
}

func NewMoviesCompilationsComposite(postgresComposite *PostgresDBComposite) (*MoviesCompilationsComposite, error) {
	MCStorage := repository.NewStorage(postgresComposite.db)
	genreStorage := genreRepository.NewStorage(postgresComposite.db)
	service := usecase.NewService(MCStorage, genreStorage)
	handler := delivery.NewHandler(service)
	return &MoviesCompilationsComposite{
		Service: service,
		Handler: handler,
	}, nil
}
