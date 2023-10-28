package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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
	user := model.User{}
	if err := s.postgres.QueryRowContext(ctx,
		`select u.id,u.password from users u where u.email = $1`,
		email).Scan(
		&user.Id,
		&user.Password,
	); err != nil {
		return "", err
	}
	fmt.Println(user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}
	token := uuid.NewString()
	expiration := time.Now().Add(time.Minute * 1)
	err := s.redis.HSet(ctx,
		token,
		[]string{
			"user",
			strconv.Itoa(int(user.Id)),
			"expiration",
			expiration.String()},
	).Err()
	if err != nil {
		return "", err
	}
	return token, nil
}

type AuthRegisterArgs struct {
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
		`insert into users(email,phone,password) values ($1,$2,$3)`,
		data.Email,
		sql.NullString{
			String: data.Phone,
			Valid:  data.Phone != "",
		},
		hashed)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			return errors.New("pg: " + pqErr.Code.Name())
		}
		return err
	}

	return nil
}

func (s Auth) Verify(ctx context.Context, token string) (bool, error) {
	return false, nil
}
