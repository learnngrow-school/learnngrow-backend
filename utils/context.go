package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDB(context *gin.Context) *gorm.DB {
	return context.MustGet("db").(*gorm.DB)
}
