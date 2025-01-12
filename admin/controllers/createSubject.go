package admin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	models "learn-n-grow.dev/m/admin/models"
	"learn-n-grow.dev/m/utils"
)

// CreateSubject Create a subject
// @summary Create a subject
// @accept  json
// @produce json
// @param   subject body models.Subject true "Subject"
// @tags    By admin
// @tags    Subjects
// @success 200 {object} string
// @router  /admin/subjects [post]
func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBind(&subject); err != nil {
		utils.Throw(c, http.StatusBadRequest, errors.New("invalid request format"))
		return
	}
	utils.GetRepo(c).CreateSubject(c, subject.Title)
	c.JSON(http.StatusOK, "ok")
}
