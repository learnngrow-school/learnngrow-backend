package auth

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.RouterGroup) {
	r.Group("/auth")
	{
		r.POST("/register", Register)
	}
}
