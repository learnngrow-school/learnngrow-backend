package reviews

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/reviews/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/reviews")
	{
		g.POST("/", reviews.Create)
		g.GET("/", reviews.GetAll)
	}
}
