package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/bio426/chocomatch-back/model"
)

type Auth struct {
	postgres *sql.DB
	redis    *redis.Client
}

func NewAuthService(pg *sql.DB, rds *redis.Client) Auth {
	service := Auth{
		postgres: pg,
		redis:    rds,
	}

	return service
}

func (s Auth) Login(ctx context.Context, email, password string) (string, error) {
	row := s.postgres.QueryRowContext(ctx, ``, email)
	user := model.User{}
	row.Scan(user.Id, user.Username)
	if err := row.Err(); err != nil {
		return "", err
	}
	token := uuid.NewString()
	err := s.redis.Set(ctx, token, "asd", 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s Auth) Register(ctx context.Context, nickname, email, phone, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	_, err = s.postgres.ExecContext(ctx,
		`insert into users(nickname,email,phone,password) values ($1,$2,$3,$4)`,
		nickname, email, phone, hashed)
	if err != nil {
		return err
	}

	return nil
}
