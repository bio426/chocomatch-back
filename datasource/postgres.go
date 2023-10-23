package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitPostgres(c Config) (*sql.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.PG_HOST, c.PG_PORT, c.PG_USER, c.PG_PASSWORD, c.PG_DATABASE)

	pg, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		return &sql.DB{}, err
	}
	if err := pg.Ping(); err != nil {
		return &sql.DB{}, err
	}
	return pg, nil
}
