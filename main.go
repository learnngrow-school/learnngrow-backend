package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"

	"learn-n-grow.dev/m/internal"
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

	v1 := r.Group("/api/v1")

	internal.AddRoutes(v1)

	r.Run()
}
