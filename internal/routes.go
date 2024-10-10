package internal

import "github.com/gin-gonic/gin"

func AddRoutes(r *gin.RouterGroup) {
	r.GET("/ping", Ping)
}
