package middlewares

import (
	"github.com/gin-gonic/gin"
	jwtUtils "learn-n-grow.dev/m/auth/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		token, err := c.Cookie("token")
		if err != nil {
			c.Set("jwtErr", err)
			c.Next()
			return
		}
		
		jwtClaims, err := jwtUtils.GetData(token)
		if err != nil {
			c.SetCookie("token", "", -1, "/", "localhost", false, true)
			c.Set("jwtErr", err)
			c.Next()
			return
		}
		email := jwtClaims["sub"].(string)

		c.Set("userEmail", email)

		c.Next()
	}
}
