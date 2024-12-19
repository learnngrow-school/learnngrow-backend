package utils

import (
	"github.com/gin-gonic/gin"
)

func Rebind[T any](c *gin.Context, input *T) *T {
	data, _ := c.Get("input")
	*input, _ = data.(T)
	return input
}
