package composites

import (
	"go.uber.org/zap"
	"myapp/internal/api"
	"myapp/internal/persons"
	"myapp/internal/persons/delivery"
	"myapp/internal/persons/repository"
	"myapp/internal/persons/usecase"
	positionsRepository "myapp/internal/position/repository"
)

type StaffComposite struct {
	Storage persons.Storage
	Service persons.Service
	Handler api.Handler
}

func NewStaffComposite(postgresComposite *PostgresDBComposite, logger *zap.SugaredLogger) (*StaffComposite, error) {
	staffStorage := repository.NewStorage(postgresComposite.db)
	personsStorage := positionsRepository.NewStorage(postgresComposite.db)
	service := usecase.NewService(staffStorage, personsStorage)
	handler := delivery.NewHandler(service, logger)
	return &StaffComposite{
		Storage: staffStorage,
		Service: service,
		Handler: handler,
	}, nil
}
