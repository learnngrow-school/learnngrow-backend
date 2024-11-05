package auth

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/auth/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	r.Group("/auth")
	{
		r.POST("/register", auth.Register)
		r.POST("/login", auth.Login)
		r.GET("/me", auth.GetMe)
	}
}
