package courses

import (
	"github.com/gin-gonic/gin"
	controllers "learn-n-grow.dev/m/courses/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/courses")
	{
		g.POST("/", controllers.CreateCourse)
		g.GET("/", controllers.GetAllCourses)
		g.GET("/:slug", controllers.GetOneCourse)
	}
}
