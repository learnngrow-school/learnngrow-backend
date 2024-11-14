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
// @router  /auth/register [post]
func Register(c *gin.Context) {
	CreateUser(c, false)
}

func CreateUser(c *gin.Context, isTeacher bool) {
	var err error
	var user auth.UserCreate

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	// TODO: check middlename type
	params := repository.CreateUserParams{IsTeacher: pgtype.Bool{Bool: isTeacher, Valid: true}}
	copier.Copy(&params, &user)

	if params.Password, err = auth.Hash(&params, user.Password); err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	_, err = internal.Server.Repo.CreateUser(context.Background(), params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	userGet := auth.UserGet{}
	copier.Copy(&userGet, &user)

	c.JSON(http.StatusCreated, userGet)
}

