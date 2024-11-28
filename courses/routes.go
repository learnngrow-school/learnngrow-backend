package courses

import (
	"github.com/gin-gonic/gin"
	models "learn-n-grow.dev/m/courses/models"
	controllers "learn-n-grow.dev/m/courses/controllers"
	"learn-n-grow.dev/m/internal/middlewares"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/courses")
	{
		g.POST("/", middlewares.Validate(models.CourseCreate{}), controllers.CreateCourse)
		g.GET("/", controllers.GetAllCourses)
		g.GET("/:id", controllers.GetOneCourse)
	}
}
