package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func New() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()

	return rdb, err
}
