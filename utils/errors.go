package utils

import "github.com/gin-gonic/gin"

func Throw(c *gin.Context, status int, err error) {
	c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}
