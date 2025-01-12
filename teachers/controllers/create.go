package teachers

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	auth "learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/utils"
)

// Create   Create teacher
// @summary Create teacher
// @accept  json
// @produce json
// @param   user body auth.UserCreate true "User"
// @tags    Teachers
// success  201 {object} int
// @router  /teachers/ [post]
func Create(c *gin.Context) {
	email, emailIsSet := c.Get("x-email")
	if !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}

	var err error

	// func CreateUser(c *gin.Context, isTeacher bool) {
	var user auth.UserCreate

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	// TODO: check middlename type
	params := repository.CreateTeacherParams{Slug: utils.GetSlug(6)}
	copier.Copy(&params, &user)

	if params.Password, err = auth.Hash(&params, user.Password); err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	id, err := utils.GetRepo(c).CreateTeacher(context.Background(), params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, id)
}
