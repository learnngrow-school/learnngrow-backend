package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	models "learn-n-grow.dev/m/admin/models"
	"learn-n-grow.dev/m/internal"
)

// CreateSubject Create a subject
// @summary Create a subject
// @produce json
// @tags    Admin
// @success 200 {object} []models.Subject
// @router  /admin/subject [get]
func GetSubject(c *gin.Context) {
	subjectDb, _ := internal.Server.Repo.GetSubjects(c)
	var subjects []models.Subject
	copier.CopyWithOption(&subjects, &subjectDb, copier.Option{DeepCopy: true})
	c.JSON(http.StatusOK, subjects)
}
