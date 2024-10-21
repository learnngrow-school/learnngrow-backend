package auth

import (
	"gorm.io/gorm"
	auth "learn-n-grow.dev/m/auth/utils"
)

type User struct {
	gorm.Model `json:"-" copier:"-"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Email      string `json:"email" gorm:"unique;not null"`
	Password   string `json:"password" gorm:"not null" copier:"-"`
}

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

func (user *User) SetHashPassword() error {
	hashed, err := auth.Hash(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return nil
}

func (user *User) CheckPassword(password string) error {
	return auth.Check(password, user.Password)
}
