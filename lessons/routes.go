package lessons

import (
	"github.com/gin-gonic/gin"
	controllers "learn-n-grow.dev/m/lessons/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/lessons")
	{
		g.POST("/", controllers.Create)
		g.GET("/:week", controllers.GetMy)
	}
}
