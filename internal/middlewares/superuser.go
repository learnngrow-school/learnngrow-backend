package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/utils"
)

func Superuser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if email, emailIsSet := c.Get("x-email"); !emailIsSet || email != "admin" {
			utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		}
	}
}
