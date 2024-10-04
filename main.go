package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func main() {
	godotenv.Load()

	r := gin.Default()
	db, _ := Connect()
	Migrate(db)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/ping", Ping)

	r.Run()
}
