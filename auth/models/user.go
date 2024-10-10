package auth

import (
	"gorm.io/gorm"
	auth "learn-n-grow.dev/m/auth/utils"
)

type User struct {
	gorm.Model `json:"-"`
	Id         uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Email      string `json:"email" gorm:"unique;not null"`
	Password   string `json:"password" gorm:"not null"`
}

func (user *User) SetHashPassword() error {
	hashed, err := auth.Hash(user.Password)

	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return nil
}
