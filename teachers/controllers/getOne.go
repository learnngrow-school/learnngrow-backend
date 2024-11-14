package teachers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"

	models "learn-n-grow.dev/m/teachers/models"
)

// GetOne   Get teacher by ID
// @summary Get teacher by ID
// @accept  json
// @produce json
// @tags    Teachers
// @param   id  path     int true "Teacher ID"
// @success 200 {object} TeacherGet
// @router  /teachers/{id} [get]
func GetOne(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.Throw(c, http.StatusUnprocessableEntity, err)
	}

	teacher, err := internal.Server.Repo.GetTeacherByID(context.Background(), int32(id))
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
	}

	var teacherRes models.TeacherGet
	copier.CopyWithOption(&teacherRes, &teacher, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, teacherRes)
}
