package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"learn-n-grow.dev/m/db"
)

// @title My magical API
// @version 1
// @description Learn & Grow API
// @basepath /api/v1
func main() {
	godotenv.Load()

	r := gin.Default()

	db.Connect()
	db.Migrate()

	r.Run()
}
