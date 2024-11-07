package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"

	"learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/internal"

	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/utils"
)

// Register Create user
// @summary Create user
// @accept  json
// @produce json
// @param   user body auth.UserCreate true "User"
// @tags    Auth
// @success 201 {object} auth.UserGet
// @router  /auth/createsuperuser [post]
func CreateSuperuser(ctx *gin.Context) {
	var user auth.UserCreate

	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.Throw(ctx, http.StatusBadRequest, err)
		return
	}

	params := repository.CreateUserParams{IsTeacher: pgtype.Bool{Bool: false, Valid: true}}
	copier.Copy(&params, &user)

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
