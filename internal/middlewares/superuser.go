package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/utils"
)

func Superuser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !checkSuperuser(c) {
			utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		}
	}
}

func checkSuperuser(c *gin.Context) bool {
	email, emailIsSet := c.Get("x-email")
	return emailIsSet && email == "admin"
}
