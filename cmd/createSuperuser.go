package cmd

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db/repository"
)

func CreateSuperuser(password string) (err error) {
	pass, err := auth.Hash([]byte(password))
	if err != nil {
		return err
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	repo := repository.New(conn)

	err = repo.CreateSuperuser(context.Background(), pass)
	return err
}
