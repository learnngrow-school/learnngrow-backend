package courses

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/courses/models"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
)

// GetOneCourse Gets one course
// @summary Get one course by id
// @accept  json
// @produce json
// @tags    Courses
// @param   slug  path     string true "Course slug"
// @success 200   {object} courses.CourseWithData
// @router  /courses/{slug}  [get]
func GetOneCourse(c *gin.Context) {
	slugParam := c.Param("slug")

	course, err := internal.Server.Repo.GetCourseBySlug(context.Background(), slugParam)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	var courseRes courses.CourseWithData
	copier.CopyWithOption(&courseRes, &course, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, courseRes)
}
