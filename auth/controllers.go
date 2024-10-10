package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/db"
	"learn-n-grow.dev/m/utils"
)

// Register Create user
// @summary Create user
// @accept json
// @produce json
// @param user body auth.UserCreate true "User"
// @tags base
// @success 200 {object} auth.UserGet
// @router /register [post]
func Register(context *gin.Context) {
	var user auth.User

	// TODO: check the binded struct type
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

	userGet := auth.UserGet{}
	copier.Copy(&userGet, &user)

	context.JSON(http.StatusCreated, userGet)
}
