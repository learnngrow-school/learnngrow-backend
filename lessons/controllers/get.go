package lessons

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/db/repository"
	models "learn-n-grow.dev/m/lessons/models"
	"learn-n-grow.dev/m/utils"
)

// GetMy Get all user's lessons
// @summary Get all user's lessons
// @tags    Lessons
// @tags    By teacher
// @tags    By student
// @accept  json
// @param   week path int true "Week"
// @produce json
// @success 200 {object} models.LessonGet[]
// @router  /lessons/{week} [get]
func GetMy(c *gin.Context) {
	var err error

	var email string
	emailAny, ok := c.Get("x-email")
	if !ok {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You have to be logged in"))
		return
	}
	email, _ = emailAny.(string)

	timeframe := GetTimeframe(c)
	params := repository.GetLessonsByUserParams{
		Email: email,
		Weeks: int32(timeframe),
	}
	lessonsDb, err := utils.GetRepo(c).GetLessonsByUser(c, params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	lessons := make([]models.LessonGet, len(lessonsDb))
	copier.CopyWithOption(&lessons, &lessonsDb, copier.Option{DeepCopy: true})

	for i := range lessons {
		lessons[i].Timestamp = int64(lessonsDb[i].Ts.Time.Unix())
	}

	c.JSON(http.StatusOK, lessons)

	// TODO: change to use cookie slug and isteacher value
	// user, err := internal.Server.Repo.GetUser(c, email)
	// if err != nil {
	// 	utils.Throw(c, http.StatusInternalServerError, err)
	// 	return
	// }
}

func GetTimeframe(c *gin.Context) int {
	stringValue := c.Param("week")

	value, _ := strconv.Atoi(stringValue)
	return value
}
