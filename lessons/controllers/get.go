package lessons

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/internal"
	"learn-n-grow.dev/m/utils"
	models "learn-n-grow.dev/m/lessons/models"
)

// GetMy Get all user's lessons
// @summary Get all user's lessons
// @tags    Lessons
// @tags    by-teacher
// @tags    by-student
// @accept  json
// @produce json
// @success 200 {object} models.LessonGet[]
// @router  /lessons/ [get]
func GetMy(c *gin.Context) {
	var err error

	var email string
	emailAny, ok := c.Get("x-email")
	if !ok {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You have to be logged in"))
		return
	}
	email, _ = emailAny.(string)

	lessonsDb, err := internal.Server.Repo.GetLessonsByTeacher(c, email)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	lessons := make([]models.LessonGet, len(lessonsDb))
	copier.CopyWithOption(&lessons, &lessonsDb, copier.Option{DeepCopy: true})

	for i := range lessons {
		lessons[i].Timestamp = 60 * int64(lessonsDb[i].TimestampM)
	}

	c.JSON(http.StatusOK, lessons)

	// TODO: change to use cookie slug and isteacher value
	// user, err := internal.Server.Repo.GetUser(c, email)
	// if err != nil {
	// 	utils.Throw(c, http.StatusInternalServerError, err)
	// 	return
	// }
	//
	// if !user.IsTeacher.Bool {
	// 	utils.Throw(c, http.StatusUnauthorized, errors.New("You have to be a teacher"))
	// 	return
	// }
}
