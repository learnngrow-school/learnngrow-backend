package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/db"
	"learn-n-grow.dev/m/utils"
)

// Register Create user
// @summary Create user
// @accept json
// @produce json
// @param user body auth.User true "User"
// @tags auth
// @router /register [post]
func Register(context *gin.Context) {
	var user auth.User

	if err := context.ShouldBindJSON(&user); err != nil {
		utils.Throw(context, http.StatusBadRequest, err)
		return
	}

	if err := user.SetHashPassword(); err != nil {
		utils.Throw(context, http.StatusInternalServerError, err)
		return
	}

	record := db.DB.Create(&user)
	if record.Error != nil {
		utils.Throw(context, http.StatusInternalServerError, record.Error)
		return
	}

	context.JSON(http.StatusCreated, user)
}
