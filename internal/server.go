package internal

import (
	"github.com/jackc/pgx/v5"
	"learn-n-grow.dev/m/db/repository"
)

type Config struct {
	Repo *repository.Queries
	Conn *pgx.Conn
	Domain string
}

var Server *Config
