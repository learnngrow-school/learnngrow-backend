package auth

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/auth/controllers"
	models "learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/internal/middlewares"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/auth")
	{
		g.POST("/register", middlewares.Validate(models.UserCreate{}), auth.Register)
		g.POST("/login", middlewares.Validate(models.UserLogin{}), auth.Login)
		g.GET("/me", auth.GetMe)
	}
}
