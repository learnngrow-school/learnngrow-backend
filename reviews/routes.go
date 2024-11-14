package reviews

import (
	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/reviews/controllers"
)

func AddRoutes(r *gin.RouterGroup) {
	g := r.Group("/")
	{
		g.POST("/reviews", reviews.Create)
		g.GET("/reviews", reviews.GetAll)
	}
}
