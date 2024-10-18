package auth

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
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
func Login(context *gin.Context) {
	var loginData auth.UserLogin

	// TODO: check the binded struct type
	if err := context.ShouldBindJSON(&loginData); err != nil {
		utils.Throw(context, http.StatusBadRequest, err)
		return
	}

	var record auth.User
	println(loginData.Email)

	err := db.DB.Model(&auth.User{}).Where("email = ?", loginData.Email).First(&record).Error

	if err != nil {
		err := errors.New("User not found")
		utils.Throw(context, http.StatusUnauthorized, err)
		return
	}

	if err := record.CheckPassword(loginData.Password); err != nil {
		utils.Throw(context, http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusAccepted, "ok")
}
