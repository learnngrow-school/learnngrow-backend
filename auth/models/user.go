package auth

import (
	auth "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db/repository"
)

type UserCreate struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"firstName" binding:"required"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserGet struct {
	ID         int32  `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
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
