package db

import (
	"github.com/garyburd/redigo/redis"
	"os"

	"github.com/joho/godotenv"
)

func ConnectRedis() (redis.Conn, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	conn, err := redis.Dial(os.Getenv("REDISPROTOCOL"), os.Getenv("REDISHOST")+":"+os.Getenv("REDISPORT"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
