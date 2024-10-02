package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", Index)

	r.Run()
}
