package datasource

import (
	"os"
)

var Config = struct {
	PG_HOST     string
	PG_PORT     string
	PG_USER     string
	PG_PASSWORD string
	PG_DATABASE string
	CLD_CLOUD   string
	CLD_KEY     string
	CLD_SECRET  string
}{}

func InitConfig() {
	Config.PG_HOST = os.Getenv("PG_HOST")
	Config.PG_PORT = os.Getenv("PG_PORT")
	Config.PG_USER = os.Getenv("PG_USER")
	Config.PG_PASSWORD = os.Getenv("PG_PASSWORD")
	Config.PG_DATABASE = os.Getenv("PG_DATABASE")

	Config.CLD_CLOUD = os.Getenv("CLD_CLOUD")
	Config.CLD_KEY = os.Getenv("CLD_KEY")
	Config.CLD_SECRET = os.Getenv("CLD_SECRET")
}
