package courses

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/internal"
)

// GetAllCourses Gets All courses
// @summary Gets All courses
// @accept  json
// @produce json
// @tags    Courses
// @success 200 {object} courses.CoursesGet
// @router  /courses/ [get]
func GetAllCourses(c *gin.Context) {
	courses, _ := internal.Server.Repo.GetAllCourses(context.Background())

	c.JSON(http.StatusOK, courses)
}
