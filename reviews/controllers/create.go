package reviews

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/db/repository"
	reviews "learn-n-grow.dev/m/reviews/models"
	"learn-n-grow.dev/m/utils"
)

// Create Create review
// @summary Create review
// @accept  json
// @produce json
// @param   user body ReviewCreate true "User"
// @tags    Reviews
// @success 201 {object} ReviewGet
// @router  /reviews/ [post]
func Create(c *gin.Context) {
	if email, emailIsSet := c.Get("x-email"); !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}

	var err error
	var reviewReq reviews.ReviewCreate
	if err = c.ShouldBindJSON(&reviewReq); err != nil {
		utils.Throw(c, http.StatusUnprocessableEntity, err)
		return
	}

	var params repository.CreateReviewParams
	copier.Copy(&params, &reviewReq)

	var queryRes repository.SchoolReview

	queryRes, err = utils.GetRepo(c).CreateReview(context.Background(), params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	var res reviews.ReviewGet

	copier.Copy(&res, &queryRes)

	c.JSON(http.StatusCreated, res)
}
