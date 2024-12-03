package auth

import (
	"gorm.io/gorm"
	auth "learn-n-grow.dev/m/auth/utils"
)

type User struct {
	gorm.Model `json:"-" copier:"-"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Email      string `json:"phone" gorm:"unique;not null"`
	Password   string `json:"password" gorm:"not null" copier:"-"`
}

type UserCreate struct {
	Email      string `json:"phone"`
	Password   string `json:"password"`
}

type UserGet struct {
	Id         uint   `json:"id"`
	Email      string `json:"phone"`
}

func (user *User) SetHashPassword() error {
	hashed, err := auth.Hash(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return nil
}
