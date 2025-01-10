package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
)

func Teacher() gin.HandlerFunc {
	return func(c *gin.Context) {
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

func checkTeacher(c *gin.Context) (bool, error) {
	var err error
	email, emailIsSet := c.Get("x-email")
	if !emailIsSet {
		return false, err
	}

	user, err := internal.Server.Repo.GetUser(c, email.(string))
	if err != nil {
		return false, err
	}

	if !user.IsTeacher.Bool {
		return false, err
	}

	return true, nil
}
