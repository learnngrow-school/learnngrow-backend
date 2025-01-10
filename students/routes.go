package students

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/internal/middlewares"
	controllers "learn-n-grow.dev/m/students/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/students")
	{
		g.GET("/", middlewares.SuperTeacher(), controllers.GetAll)
	}
}
