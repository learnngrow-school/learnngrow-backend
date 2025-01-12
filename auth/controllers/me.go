package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/utils"
)

// GetMe    Get current user data
// @summary Get current user data
// @accept  json
// @produce json
// @tags    Auth
// @success 200 {object} auth.UserMe
// @router  /auth/me [get]
func GetMe(c *gin.Context) {
	var err error

	email, jwtIsSet := c.Get("x-email")
	if !jwtIsSet {
		err, hasErr := c.Get("x-jwt-err")
		if !hasErr {
			panic("\n\nNo cookie but no jwt error set\n")
		}
		utils.Throw(c, http.StatusUnauthorized, err.(error))
		return
	}

	var record repository.User
	record, err = utils.GetRepo(c).GetUser(context.Background(), email.(string))
	if err != nil {
		err := errors.New("No such user found")
		utils.Throw(c, http.StatusUnauthorized, err)
		return
	}

	userMe := auth.UserMe{}
	copier.Copy(&userMe, &record)

	c.JSON(http.StatusOK, userMe)
}
