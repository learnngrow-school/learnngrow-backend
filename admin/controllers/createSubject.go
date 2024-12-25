package admin

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	models "learn-n-grow.dev/m/admin/models"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
)

// CreateSubject Create a subject
// @summary Create a subject
// @accept  json
// @produce json
// @param   subject body models.Subject true "Subject"
// @tags    Admin
// @success 200 {object} string
// @router  /admin/subject [post]
func CreateSubject(c *gin.Context) {
	var subject models.Subject
	if err := c.ShouldBind(&subject); err != nil {
		utils.Throw(c, http.StatusBadRequest, errors.New("invalid request format"))
	}
	internal.Server.Repo.CreateSubject(c, subject.Title)
	c.JSON(http.StatusOK, "ok")
}
