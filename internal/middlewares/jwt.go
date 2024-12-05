package middlewares

import (
	"github.com/gin-gonic/gin"
	jwtUtils "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/internal"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		token, err := c.Cookie("token")
		if err != nil {
			c.Set("x-jwt-err", err)
			c.Next()
			return
		}
		
		jwtClaims, err := jwtUtils.GetData(token)
		if err != nil {
			c.SetCookie("token", "", -1, "/", internal.Server.Domain, false, true)
			c.Set("x-jwt-err", err)
			c.Next()
			return
		}
		email := jwtClaims["sub"].(string)

		c.Set("x-email", email)

		c.Next()
	}
}
