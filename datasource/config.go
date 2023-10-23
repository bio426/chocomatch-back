package datasource

import (
	"os"
)

type Config = struct {
	PG_HOST      string
	PG_PORT      string
	PG_USER      string
	PG_PASSWORD  string
	PG_DATABASE  string
	RDS_ADDR     string
	RDS_USER     string
	RDS_PASSWORD string
	CLD_CLOUD    string
	CLD_KEY      string
	CLD_SECRET   string
}

func InitConfig() Config {
	result := Config{}

	result.PG_HOST = os.Getenv("PG_HOST")
	result.PG_PORT = os.Getenv("PG_PORT")
	result.PG_USER = os.Getenv("PG_USER")
	result.PG_PASSWORD = os.Getenv("PG_PASSWORD")
	result.PG_DATABASE = os.Getenv("PG_DATABASE")

	result.RDS_ADDR = os.Getenv("RDS_ADDR")
	result.RDS_USER = os.Getenv("RDS_USER")
	result.RDS_PASSWORD = os.Getenv("RDS_PASSWORD")

	result.CLD_CLOUD = os.Getenv("CLD_CLOUD")
	result.CLD_KEY = os.Getenv("CLD_KEY")
	result.CLD_SECRET = os.Getenv("CLD_SECRET")

	return result
}
