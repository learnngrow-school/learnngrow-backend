package reviews

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/internal"
	reviews "learn-n-grow.dev/m/reviews/models"
	"learn-n-grow.dev/m/utils"
)

type school struct {}
type teacher struct {}
type course struct {}

// Create Create review
// @summary Create review
// @accept  json
// @produce json
// @param   review body ReviewCreate true "User"
// @tags    Reviews
// @success 201 {object} ReviewGet
// @router  /school/reviews [post]
func Create[T any](c *gin.Context) {
	var err error

	if email, emailIsSet := c.Get("userEmail"); !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}

	var x T
	switch any(x).(type) {
	case school:
		reviewReq := reviews.ReviewCreate{}
		params := repository.CreateSchoolReviewParams{}
		queryRes := repository.SchoolReview{}
		create := internal.Server.Repo.CreateSchoolReview
		Hacer(&reviewReq, &params, &queryRes, &create)
	}


	reviewReq := *utils.Rebind(c, &reviews.ReviewCreate{})

	var params repository.CreateSchoolReviewParams
	copier.Copy(&params, &reviewReq)

	var queryRes repository.SchoolReview
	queryRes, err = internal.Server.Repo.CreateSchoolReview(context.Background(), params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
	}

	var res reviews.ReviewGet
	copier.Copy(&res, &queryRes)

	c.JSON(http.StatusCreated, res)
}

func Hacer[T any]() {

}

// Create Create teacher review
// @summary Create teacher review
// @accept  json
// @produce json
// @param   id   path int                 true "Teacher Id"
// @param   user body TeacherReviewCreate true "Review"
// @tags    Reviews, Teachers
// @success 201 {object} TeacherReviewGet
// @router  /teachers/{id}/reviews [post]
func TeacherCreate(c *gin.Context) {
	var err error

	if email, emailIsSet := c.Get("userEmail"); !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}

	reviewReq := *utils.Rebind(c, &reviews.ReviewCreate{})

	var params repository.CreateSchoolReviewParams
	copier.Copy(&params, &reviewReq)

	var queryRes repository.SchoolReview
	queryRes, err = internal.Server.Repo.CreateSchoolReview(context.Background(), params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
	}

	var res reviews.ReviewGet
	copier.Copy(&res, &queryRes)

	c.JSON(http.StatusCreated, res)
}


// Create Create course review
// @summary Create course review
// @accept  json
// @produce json
// @param   id   path int                true "Course Id"
// @param   user body CourseReviewCreate true "Review"
// @tags    Reviews, Courses
// @success 201 {object} CourseReviewGet
// @router  /courses/{id}/reviews [post]
func CourseCreate(c *gin.Context) {
	var err error

	if email, emailIsSet := c.Get("userEmail"); !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}

	reviewReq := *utils.Rebind(c, &reviews.ReviewCreate{})

	var params repository.CreateSchoolReviewParams
	copier.Copy(&params, &reviewReq)

	var queryRes repository.SchoolReview
	queryRes, err = internal.Server.Repo.CreateSchoolReview(context.Background(), params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
	}

	var res reviews.ReviewGet
	copier.Copy(&res, &queryRes)

	c.JSON(http.StatusCreated, res)
}

