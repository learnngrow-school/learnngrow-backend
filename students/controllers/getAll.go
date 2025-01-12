package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	auth "learn-n-grow.dev/m/auth/models"
	"learn-n-grow.dev/m/utils"
)

// GetAll   Get all students
// @summary Get all students
// @produce json
// @tags    Students
// @tags    _ By admin
// @tags    _ By teacher
// @success 200 {object} []auth.UserGet
// @router  /students/ [get]
func GetAll(c *gin.Context) {
	studentsDb, err := utils.GetRepo(c).GetAllStudents(c)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	var students []auth.UserGet
	copier.CopyWithOption(&students, &studentsDb, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, students)
}
