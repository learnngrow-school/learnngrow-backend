package reviews

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/internal"
	reviews "learn-n-grow.dev/m/reviews/models"
	"learn-n-grow.dev/m/utils"
)

// GetAll Get all school reviews
// @summary Get all school reviews
// @accept  json
// @produce json
// @tags    Reviews
// @success 200 {object} []reviews.ReviewGet
// @router  /reviews/ [get]
func GetAll(c *gin.Context) {
	var err error
	var queryRes []repository.SchoolReview
	queryRes, err = internal.Server.Repo.GetAllSchoolReviews(context.Background())
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
	}

	var res []reviews.ReviewGet
	copier.CopyWithOption(&res, &queryRes, copier.Option{DeepCopy: true})

	c.JSON(http.StatusOK, res)
}
