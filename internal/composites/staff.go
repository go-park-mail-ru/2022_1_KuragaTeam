package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/db/position"
	"myapp/internal/persons"
	"myapp/internal/persons/delivery"
	"myapp/internal/persons/repository"
	"myapp/internal/persons/usecase"
)

type StaffComposite struct {
	Storage persons.Storage
	Service persons.Service
	Handler api.Handler
}

func NewStaffComposite(postgresComposite *PostgresDBComposite) (*StaffComposite, error) {
	staffStorage := repository.NewStorage(postgresComposite.db)
	positionStorage := position.NewStorage(postgresComposite.db)
	service := usecase.NewService(staffStorage, positionStorage)
	handler := delivery.NewHandler(service)
	return &StaffComposite{
		Storage: staffStorage,
		Service: service,
		Handler: handler,
	}, nil
}
