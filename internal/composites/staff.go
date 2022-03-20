package composites

import (
	"myapp/internal/adapters/api"
	staff2 "myapp/internal/adapters/api/staff"
	position2 "myapp/internal/adapters/db/position"
	staff3 "myapp/internal/adapters/db/staff"
	staff "myapp/internal/domain/staff"
)

type StaffComposite struct {
	Storage staff.Storage
	Service staff2.Service
	Handler api.Handler
}

func NewStaffComposite(postgresComposite *PostgresDBComposite) (*StaffComposite, error) {
	staffStorage := staff3.NewStorage(postgresComposite.db)
	positionStorage := position2.NewStorage(postgresComposite.db)
	service := staff.NewService(staffStorage, positionStorage)
	handler := staff2.NewHandler(service)
	return &StaffComposite{
		Storage: staffStorage,
		Service: service,
		Handler: handler,
	}, nil
}
