package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   //`json:"id" gorm:"primaryKey;autoIncrement:true"`
	Email    string //`json:"email" gorm:"unique;not null"`
	Password string //`json:"password" gorm:"not null"`
}
