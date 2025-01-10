package courses

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"
	models "learn-n-grow.dev/m/courses/models"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
)

// CreateCourse Create a course
// @summary Create a course
// @accept  json
// @produce json
// @param   user body courses.CourseCreate true "Course"
// @tags    Courses
// @success 201 {object} courses.Course
// @router  /courses/ [post]
func CreateCourse(c *gin.Context) {
	email, emailIsSet := c.Get("x-email")
	if !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}

	var err error
	var course models.CourseCreate

	if err = c.ShouldBindJSON(&course); err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	params := *getCreateCourseParams(&course)
	record, err := internal.Server.Repo.CreateCourse(context.Background(), params)

	c.JSON(http.StatusCreated, record)
}

func getCreateCourseParams(course *models.CourseCreate) *repository.CreateCourseParams {
	params := repository.CreateCourseParams{
		Year: pgtype.Int2{
			Int16: course.Year,
			Valid: true,
		},
		Description: pgtype.Text{
			String: course.Description,
			Valid:  course.Description != "",
		},
		Slug: utils.GetSlug(6),
	}
	copier.Copy(&params, &course)
	return &params
}
