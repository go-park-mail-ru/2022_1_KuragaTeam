package composites

import (
	"go.uber.org/zap"
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

func NewMoviesCompilationsComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*MoviesCompilationsComposite, error) {
	MCStorage := repository.NewStorage(postgresComposite.Db)
	genreStorage := genreRepository.NewStorage(postgresComposite.Db)
	service := usecase.NewService(MCStorage, genreStorage)
	handler := delivery.NewHandler(service, logger)
	return &MoviesCompilationsComposite{
		Service: service,
		Handler: handler,
	}, nil
}
