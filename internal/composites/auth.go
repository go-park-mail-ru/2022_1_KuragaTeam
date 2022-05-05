package composites

import (
	"myapp/internal/microservices/authorization"
	"myapp/internal/microservices/authorization/repository"
	"myapp/internal/microservices/authorization/usecase"
)

type AuthComposite struct {
	Storage authorization.Storage
	Service *usecase.Service
}

func NewAuthComposite(postgresComposite *PostgresDBComposite, redisComposite *RedisComposite) (*AuthComposite, error) {
	storage := repository.NewStorage(postgresComposite.db, redisComposite.redis)
	service := usecase.NewService(storage)
	return &AuthComposite{
		Storage: storage,
		Service: service,
	}, nil
}
