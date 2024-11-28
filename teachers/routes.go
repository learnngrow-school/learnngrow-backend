package teachers

import (
	"github.com/gin-gonic/gin"
	auth "learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/internal/middlewares"
	"learn-n-grow.dev/m/teachers/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/teachers")
	{
		g.POST("/", middlewares.Validate(auth.UserCreate{}), teachers.Create)
		g.GET("/", teachers.GetAll)
		g.GET("/:id", teachers.GetOne)
	}
}
