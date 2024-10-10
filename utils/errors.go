package utils

import "github.com/gin-gonic/gin"

func Throw(context *gin.Context, status int, err error) {
	context.JSON(status, gin.H{"error": err.Error()})
	context.Abort()
}
