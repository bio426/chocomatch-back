package service

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type Auth struct {
	postgres *sql.DB
	redis    *redis.Client
}

func (s Auth) Login() error {
	return nil
}

func NewAuthService(pg *sql.DB, rds *redis.Client) Auth {
	service := Auth{
		postgres: pg,
		redis:    rds,
	}

	return service
}
