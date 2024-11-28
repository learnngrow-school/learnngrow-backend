package auth

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
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

func (user UserCreate) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.E164),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
		validation.Field(&user.FirstName, validation.Required, validation.Length(3, 100)),
		validation.Field(&user.MiddleName, validation.Length(3, 100)),
		validation.Field(&user.LastName, validation.Required, validation.Length(3, 100)),
	)
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user UserLogin) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
	)
}

type UserGet struct {
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
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
