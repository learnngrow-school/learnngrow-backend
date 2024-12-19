package reviews

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/internal/middlewares"
	models "learn-n-grow.dev/m/reviews/models"
	"learn-n-grow.dev/m/reviews/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/")
	{
		g.POST("/school/reviews", middlewares.Validate(models.ReviewCreate{}), reviews.Create)
		g.GET("/school/reviews", reviews.GetAll)
		g.POST("/teachers/:id/reviews", middlewares.Validate(models.ReviewCreate{}), reviews.TeacherCreate)
		g.GET("/teachers/:id/reviews", reviews.TeacherGetAll)
		g.POST("/courses/:id/reviews", middlewares.Validate(models.ReviewCreate{}), reviews.CourseCreate)
		g.GET("/courses/:id/reviews", reviews.CourseGetAll)
	}
}
