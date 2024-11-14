package courses

import (
	"context"
	"net/http"
	"strconv"

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
// @param   id  path     int true "Course ID"
// @success 200 {object} courses.CourseWithData
// @router  /courses/{id} [get]
func GetOneCourse(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Throw(c, http.StatusUnprocessableEntity, err)
	}

	course, _ := internal.Server.Repo.GetCourseById(context.Background(), int32(id))

	var courseRes courses.CourseWithData

	copier.Copy(&courseRes, &course)

	c.JSON(http.StatusOK, courseRes)
}
