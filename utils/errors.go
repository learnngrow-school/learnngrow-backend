package utils

import "github.com/gin-gonic/gin"

func Throw(context *gin.Context, status int, err error) {
	context.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}
