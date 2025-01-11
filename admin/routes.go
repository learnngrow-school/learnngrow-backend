package admin

import (
	"github.com/gin-gonic/gin"
	controllers "learn-n-grow.dev/m/admin/controllers"
	"learn-n-grow.dev/m/internal/middlewares"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/admin")
	{
		g.POST("/subjects", middlewares.Superuser(), controllers.CreateSubject)
		g.GET("/subjects", controllers.GetSubjects)
		g.POST("/files", middlewares.SuperTeacher(), controllers.UploadFile)
	}
}
