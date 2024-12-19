package admin

import (
	"github.com/gin-gonic/gin"
	controllers "learn-n-grow.dev/m/admin/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/")
	{
		g.POST("/subjects", controllers.Create)
	}
}
