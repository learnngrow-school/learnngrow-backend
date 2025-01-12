package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"learn-n-grow.dev/m/admin"
	"learn-n-grow.dev/m/cmd"
	"learn-n-grow.dev/m/courses"
	_ "learn-n-grow.dev/m/docs"
	"learn-n-grow.dev/m/internal/middlewares"
	"learn-n-grow.dev/m/lessons"
	"learn-n-grow.dev/m/reviews"
	"learn-n-grow.dev/m/students"
	"learn-n-grow.dev/m/teachers"
	"learn-n-grow.dev/m/utils"

	"learn-n-grow.dev/m/auth"
	"learn-n-grow.dev/m/internal"
)

func startServer() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.JWTMiddleware())
	r.Use(middlewares.DbMiddleware())
	r.Use(middlewares.LoggerMiddleware())
	utils.SetupSlugs()

	v1 := r.Group("/api/v1")

	internal.AddRoutes(v1)
	auth.AddRoutes(v1)
	teachers.AddRoutes(v1)
	courses.AddRoutes(v1)
	reviews.AddRoutes(v1)
	lessons.AddRoutes(v1)
	admin.AddRoutes(v1)
	students.AddRoutes(v1)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})

	r.Run(":8000")
}

// @title My magical API
// @version 1
// @description Learn & Grow API
// @basepath /api/v1
func main() {
	godotenv.Load()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	internal.Server = &internal.Config{
		Domain: "localhost",
	}

	if len(os.Args) > 1 && os.Args[1] == "createsuperuser" {
		if err := cmd.CreateSuperuser(os.Args[2]); err != nil {
			fmt.Println(err)
		}
		return
	}

	startServer()
}
