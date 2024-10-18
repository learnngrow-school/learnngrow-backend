package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
	jwtUtil "learn-n-grow.dev/m/auth/utils"
	"learn-n-grow.dev/m/db"
	"learn-n-grow.dev/m/utils"
)

// Login Log in user, issue JWT
// @summary Log in user, issue JWT into Cookie
// @accept json
// @produce json
// @param user body auth.UserLogin true "User"
// @tags base
// @success 200 {object} auth.UserGet
// @router /login [post]
func Login(c *gin.Context) {
	var loginData auth.UserLogin

	// TODO: check the binded struct type
	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	var record auth.User

	err := db.DB.Model(&auth.User{}).Where("email = ?", loginData.Email).First(&record).Error

	if err != nil {
		err := errors.New("User not found")
		utils.Throw(c, http.StatusUnauthorized, err)
		return
	}

	if err := record.CheckPassword(loginData.Password); err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	println(record.Email)
	jwt, err := jwtUtil.Issue(record.Email)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("token", jwt, int(jwtUtil.ExpTime), "/", "localhost", false, true)
	c.JSON(http.StatusAccepted, "ok")
}
