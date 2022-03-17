package composites

import (
	"myapp/internal/adapters/api"
	"myapp/internal/adapters/api/user"
	user3 "myapp/internal/adapters/db/user"
	user2 "myapp/internal/domain/user"
)

type UserComposite struct {
	Storage    user2.Storage
	Service    user.Service
	Handler    api.Handler
	Middleware api.Middleware
}

func NewUserComposite(postgresComposite *PostgresDBComposite, redisComposite *RedisComposite) (*UserComposite, error) {
	storage := user3.NewStorage(postgresComposite.db)
	redis := user3.NewRedisStore(redisComposite.redis)
	service := user2.NewService(storage, redis)
	handler := user.NewHandler(service)
	middleware := user.NewMiddleware(service)
	return &UserComposite{
		Storage:    storage,
		Service:    service,
		Handler:    handler,
		Middleware: middleware,
	}, nil
}
