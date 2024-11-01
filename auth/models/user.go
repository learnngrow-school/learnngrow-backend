package auth

import (
	auth "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db/repository"
)

type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserGet struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func SetHashPassword(params *repository.CreateUserParams, password string) error {
	hashed, err := auth.Hash([]byte(password))

	if err != nil {
		return err
	}

	params.Password = hashed
	return nil
}

func CheckPassword(user repository.User, password string) error {
	return auth.Check(password, user.Password)
}
