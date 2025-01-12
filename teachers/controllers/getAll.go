package teachers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/utils"

	models "learn-n-grow.dev/m/teachers/models"
)

// GetAll   Get all teachers
// @summary Get all teachers
// @accept  json
// @produce json
// @tags    Teachers
// @success 200 {object} TeachersGet
// @router  /teachers/ [get]
func GetAll(c *gin.Context) {
	teachers, err := utils.GetRepo(c).GetAllTeachers(context.Background())
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	var teachersRes models.TeachersGet
	copier.CopyWithOption(&teachersRes, &teachers, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, teachersRes)
}
