package datasource

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Postgres *sql.DB

func InitPostgres() (*sql.DB, error) {
	connectionInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable")

	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
}
