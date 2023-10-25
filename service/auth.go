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
	row := s.postgres.QueryRowContext(ctx, `select u.id,u.password from users u where u.email = $1`, email)
	user := model.User{}
	row.Scan(user.Id, user.Password)
	if err := row.Err(); err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token := uuid.NewString()
	err := s.redis.Set(ctx, token, "asd", 0).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

type AuthRegisterArgs struct {
	Username string
	Email    string
	Phone    string
	Password string
}

func (s Auth) Register(ctx context.Context, data AuthRegisterArgs) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return err
	}
	_, err = s.postgres.ExecContext(ctx,
		`insert into users(username,email,phone,password) values ($1,$2,$3,$4)`,
		data.Username, data.Email, data.Phone, hashed)
	if err != nil {
		return err
	}

	return nil
}
