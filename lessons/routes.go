package lessons

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/internal/middlewares"
	controllers "learn-n-grow.dev/m/lessons/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/lessons")
	{
		g.POST("/", middlewares.SuperTeacher(), controllers.Create)
		g.GET("/:week", controllers.GetMy)
	}
}
