package courses

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/courses/models"
	"learn-n-grow.dev/m/utils"
)

// GetAllCourses Gets All courses
// @summary Gets All courses
// @accept  json
// @produce json
// @tags    Courses
// @success 200 {object} courses.CoursesGet
// @router  /courses/ [get]
func GetAllCourses(c *gin.Context) {
	coursesQuery, _ := utils.GetRepo(c).GetAllCourses(context.Background())

	var coursesRes courses.CoursesGet

	copier.CopyWithOption(&coursesRes, &coursesQuery, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, coursesRes)
}
