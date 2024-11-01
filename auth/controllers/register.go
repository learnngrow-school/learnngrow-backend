package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/internal"

	"learn-n-grow.dev/m/db/repository"
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
func Register(ctx *gin.Context) {
	var user auth.UserCreate

	// TODO: check the binded struct type
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.Throw(ctx, http.StatusBadRequest, err)
		return
	}

	params := repository.CreateUserParams{ Email: user.Email }

	if err := auth.SetHashPassword(&params, user.Password); err != nil {
		utils.Throw(ctx, http.StatusInternalServerError, err)
		return
	}

	_, err := internal.Server.Repo.CreateUser(context.Background(), params)
	if err != nil {
		utils.Throw(ctx, http.StatusInternalServerError, err)
		return
	}

	userGet := auth.UserGet{}
	copier.Copy(&userGet, &user)

	ctx.JSON(http.StatusCreated, userGet)
}
