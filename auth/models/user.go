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
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Slug       string `json:"slug"`
}

func Hash[T repository.CreateUserParams | repository.CreateTeacherParams] (params *T, password string) (hashed []byte, err error) {
	hashed, err = auth.Hash([]byte(password))

	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func CheckPassword(user repository.User, password string) error {
	return auth.Check(password, user.Password)
}
