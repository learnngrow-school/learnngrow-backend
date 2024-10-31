package auth

import (
	// "gorm.io/gorm"
	auth "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db/repository"
)

// type User struct {
// 	gorm.Model `json:"-" copier:"-"`
// 	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
// 	Email      string `json:"email" gorm:"unique;not null"`
// 	Password   string `json:"password" gorm:"not null" copier:"-"`
// }

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
