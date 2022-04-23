package composites

import (
	api2 "myapp/internal/api"
	"myapp/internal/user"
	"myapp/internal/user/delivery"
	"myapp/internal/user/repository"
	"myapp/internal/user/usecase"

	"go.uber.org/zap"
)

type UserComposite struct {
	Storage user.Storage
	Service user.Service
	Handler api2.Handler
}

func NewUserComposite(postgresComposite *PostgresDBComposite, minioComposite *MinioComposite, logger *zap.SugaredLogger) (*UserComposite, error) {
	storage := repository.NewStorage(postgresComposite.db)
	minio := repository.NewImageStorage(minioComposite.client)
	service := usecase.NewService(storage, minio)
	handler := delivery.NewHandler(service, logger)
	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
