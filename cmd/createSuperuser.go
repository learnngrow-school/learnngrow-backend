package cmd

import (
	"context"

	"learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/internal"
)

func CreateSuperuser(password string) (err error) {
	pass, err := auth.Hash([]byte(password))
	if err != nil {
		return err
	}

	err = internal.Server.Repo.CreateSuperuser(context.Background(), pass)
	return err
}
