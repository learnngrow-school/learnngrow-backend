package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/utils"
)

func SuperTeacher() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin := checkSuperuser(c)
		if isAdmin {
			return
		}

		isTeacher, err := checkTeacher(c)
		if err != nil {
			utils.Throw(c, http.StatusInternalServerError, err)
			return
		}
		if !isTeacher {
			utils.Throw(c, http.StatusUnauthorized, errors.New("You are not a teacher"))
			return
		}
	}
}
