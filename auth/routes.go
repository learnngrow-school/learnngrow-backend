package auth

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/auth/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/auth")
	{
		g.POST("/register", auth.Register)
		g.POST("/login", auth.Login)
		g.GET("/me", auth.GetMe)
	}
}
