package composites

import (
	"myapp/internal/microservices/profile"
	"myapp/internal/microservices/profile/repository"
	"myapp/internal/microservices/profile/usecase"
)

type ProfileComposite struct {
	Storage profile.Storage
	Service *usecase.Service
}

func NewProfileComposite(postgresComposite *PostgresDBComposite, minioComposite *MinioComposite,
	redisComposite *RedisComposite) (*ProfileComposite, error) {
	storage := repository.NewStorage(postgresComposite.db, minioComposite.client, redisComposite.redis)
	service := usecase.NewService(storage)
	return &ProfileComposite{
		Storage: storage,
		Service: service,
	}, nil
}
