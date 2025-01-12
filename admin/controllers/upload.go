package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/utils"
)

type Slug struct {
	Slug string `json:"slug"`
}

// UploadFile Upload a file
// @summary   Upload a file
// @accept    mpfd
// @produce   json
// @param     file formData file true "File"
// @tags      Files
// @tags      By admin
// @tags      By teacher
// @success   201 {object} Slug
// @router    /admin/files [post]
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	slug, err := utils.Upload(c, file)
	if err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, Slug{Slug: slug})
}
