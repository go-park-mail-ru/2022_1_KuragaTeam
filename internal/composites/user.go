package composites

import (
	api2 "myapp/internal/api"
	"myapp/internal/middleware"
	"myapp/internal/user"
	"myapp/internal/user/delivery"
	"myapp/internal/user/repository"
	"myapp/internal/user/usecase"
)

type UserComposite struct {
	Storage    user.Storage
	Service    user.Service
	Handler    api2.Handler
	Middleware api2.Middleware
}

func NewUserComposite(postgresComposite *PostgresDBComposite, redisComposite *RedisComposite, minioComposite *MinioComposite) (*UserComposite, error) {
	storage := repository.NewStorage(postgresComposite.db)
	redis := repository.NewRedisStore(redisComposite.redis)
	minio := repository.NewImageStorage(minioComposite.client)
	service := usecase.NewService(storage, redis, minio)
	handler := delivery.NewHandler(service)
	middleware := middleware.NewMiddleware(service)
	return &UserComposite{
		Storage:    storage,
		Service:    service,
		Handler:    handler,
		Middleware: middleware,
	}, nil
}
