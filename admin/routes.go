package admin

import (
	"github.com/gin-gonic/gin"
	controllers "learn-n-grow.dev/m/admin/controllers"
	"learn-n-grow.dev/m/internal/middlewares"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/admin")
	{
		g.POST("/subject", middlewares.Superuser(), controllers.CreateSubject)
		g.GET("/subject", controllers.GetSubject)
	}
}
