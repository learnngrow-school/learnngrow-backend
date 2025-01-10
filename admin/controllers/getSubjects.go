package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	models "learn-n-grow.dev/m/admin/models"
	"learn-n-grow.dev/m/internal"
)

// CreateSubject Get list of subjects
// @summary Get list of subjects
// @produce json
// @tags    Subjects
// @success 200 {object} []models.Subject
// @router  /admin/subjects [get]
func GetSubjects(c *gin.Context) {
	subjectDb, _ := internal.Server.Repo.GetSubjects(c)
	var subjects []models.Subject
	copier.CopyWithOption(&subjects, &subjectDb, copier.Option{DeepCopy: true})
	c.JSON(http.StatusOK, subjects)
}
