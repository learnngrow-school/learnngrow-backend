package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	auth "learn-n-grow.dev/m/auth/models"
	jwtUtil "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
)

// Login    Log in user, issue JWT
// @summary Log in user, issue JWT into Cookie
// @accept  json
// @produce json
// @param   user body auth.UserLogin true "User"
// @tags    Auth
// @success 202 {object} auth.UserGet
// @router  /auth/login [post]
func Login(c *gin.Context) {
	data, _ := c.Get("input")
	loginData, _ := data.(auth.UserLogin)

	var record repository.User
	record, err := internal.Server.Repo.GetUser(context.Background(), loginData.Email)
	if err != nil {
		err := errors.New("User not found")
		utils.Throw(c, http.StatusUnauthorized, err)
		return
	}

	if err := auth.CheckPassword(record, loginData.Password); err != nil {
		utils.Throw(c,
			http.StatusInternalServerError,
			errors.New("Wrong password"),
		)
		return
	}

	jwt, err := jwtUtil.Issue(record.Email)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	userGet := auth.UserGet{}
	copier.Copy(&userGet, &record)

	c.SetCookie("token", jwt, int(jwtUtil.ExpTime), "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, userGet)
}
