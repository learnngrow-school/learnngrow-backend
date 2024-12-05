package teachers

import (
	"context"
	"net/http"

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
// @param   slug  path       string true "Teacher Slug"
// @success 200   {object}   TeacherGet
// @router  /teachers/{slug} [get]
func GetOne(c *gin.Context) {
	slugParam := c.Param("slug")
	// id, err := strconv.Atoi(idParam)
	// if err != nil {
	// 	utils.Throw(c, http.StatusUnprocessableEntity, err)
	// }

	teacher, err := internal.Server.Repo.GetTeacherBySlug(context.Background(), slugParam)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
	}

	var teacherRes models.TeacherGet
	copier.CopyWithOption(&teacherRes, &teacher, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, teacherRes)
}
