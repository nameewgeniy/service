package repository

import (
	"github.com/jackc/pgx/v4"
)

type ConfigRepository struct {
	DB *pgx.Conn
}
