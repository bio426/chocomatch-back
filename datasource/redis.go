package datasource

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func InitRedis(c Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.RDS_ADDR,
		Username: c.RDS_USER,
		Password: c.RDS_PASSWORD,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return &redis.Client{}, err
	}

	return rdb, nil
}
