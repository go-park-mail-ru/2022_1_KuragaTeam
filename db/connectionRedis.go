package db

import (
	"os"
	"time"

	"github.com/gomodule/redigo/redis"

	"github.com/joho/godotenv"
)

func ConnectRedis() (*redis.Pool, error) {
	if err := godotenv.Load(".env"); err != nil {
		return &redis.Pool{}, err
	}

	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) {
			return redis.Dial(os.Getenv("REDISPROTOCOL"), os.Getenv("REDISHOST")+":"+os.Getenv("REDISPORT"))
		},
	}, nil
}
