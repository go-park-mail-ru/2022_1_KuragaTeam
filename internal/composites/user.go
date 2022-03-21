package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/user"
	"myapp/internal/user/delivery"
	"myapp/internal/user/repository"
	"myapp/internal/user/usecase"
)

type UserComposite struct {
	Storage    user.Storage
	Service    user.Service
	Handler    api.Handler
	Middleware api.Middleware
}

func NewUserComposite(postgresComposite *PostgresDBComposite, redisComposite *RedisComposite) (*UserComposite, error) {
	storage := repository.NewStorage(postgresComposite.db)
	redis := repository.NewRedisStore(redisComposite.redis)
	service := usecase.NewService(storage, redis)
	handler := delivery.NewHandler(service)
	middleware := user.NewMiddleware(service)
	return &UserComposite{
		Storage:    storage,
		Service:    service,
		Handler:    handler,
		Middleware: middleware,
	}, nil
}
