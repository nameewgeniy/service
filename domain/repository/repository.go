package repository

import (
	"github.com/jackc/pgx/v4"
)

type ConfigRepository struct {
	DB *pgx.Conn
}

var config *ConfigRepository

func InitConfig(cf *ConfigRepository)  {
	config = cf
}