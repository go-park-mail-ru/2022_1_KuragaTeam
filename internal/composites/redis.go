package composites

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisComposite struct {
	redis *redis.Pool
}

func NewRedisComposite() (*RedisComposite, error) {
	redisPool := redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(os.Getenv("REDISPROTOCOL"), os.Getenv("REDISHOST")+":"+os.Getenv("REDISPORT"))
		},
	}

	return &RedisComposite{redis: &redisPool}, nil
}
