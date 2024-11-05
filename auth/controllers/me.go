package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
	jwtUtil "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
)

// Get current user data
// @summary Get current user data
// @accept json
// @produce json
// @tags base
// @success 200 {object} auth.UserGet
// @router /me [get]
func GetMe(c *gin.Context) {
	var err error

	token, err := c.Cookie("token")
	if err != nil {
		utils.Throw(c, http.StatusUnauthorized, err)
	}
	
	jwtClaims, err := jwtUtil.GetData(token)
	if err != nil {
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		utils.Throw(c, http.StatusUnauthorized, err)
		return
	}
	email := jwtClaims["sub"].(string)

	var record repository.User
	record, err = internal.Server.Repo.GetUser(context.Background(), email)
	if err != nil {
		err := errors.New("No such user found")
		utils.Throw(c, http.StatusUnauthorized, err)
		return
	}

	userGet := auth.UserGet{}
	copier.Copy(&userGet, &record)

	c.JSON(http.StatusOK, userGet)
}
