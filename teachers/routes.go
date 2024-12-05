package teachers

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/teachers/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/teachers")
	{
		g.POST("/", teachers.Create)
		g.GET("/", teachers.GetAll)
		g.GET("/:slug", teachers.GetOne)
	}
}
