package internal

import (
	"learn-n-grow.dev/m/db/repository"
)

type Config struct {
	Repo *repository.Queries
}

var Server *Config

