package lessons

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/db/repository"
	"learn-n-grow.dev/m/internal"
	models "learn-n-grow.dev/m/lessons/models"
	"learn-n-grow.dev/m/utils"
)

// Create Create a lesson
// @summary Create a lesson
// @accept  json
// @produce json
// @param   user body models.LessonCreate true "User"
// @tags    Lessons
// @tags    by teacher
// @success 201 {object} models.LessonGet
// @router  /lessons/ [post]
func Create(c *gin.Context) {
	var err error

	var email string
	emailAny, ok := c.Get("x-email")
	if !ok {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You have to be logged in"))
		return
	}
	email, _ = emailAny.(string)

	teacher, err := internal.Server.Repo.GetUser(c, email)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	if !teacher.IsTeacher.Bool {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You have to be a teacher"))
		return
	}

	var lesson models.LessonCreate
	if err = c.ShouldBindJSON(&lesson); err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	params := GetCreateLessonParams(lesson, teacher)
	result, err := internal.Server.Repo.CreateLesson(c, params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	var lessonResult models.LessonGet
	copier.CopyWithOption(&lessonResult, &result, copier.Option{DeepCopy: true})
	lessonResult.Timestamp = lesson.Timestamp

	c.JSON(http.StatusCreated, lessonResult)
}

func GetCreateLessonParams(lesson models.LessonCreate, teacher repository.User) repository.CreateLessonParams {
	// for maybe allowing superuser in the future
	teacherSlug := teacher.Slug
	return repository.CreateLessonParams{
		Slug:         lesson.StudentSlug,
		Slug_2:       teacherSlug,
		Ts:           pgtype.Timestamptz{Time: time.Unix(lesson.Timestamp, 0), Valid: true},
		TeacherNotes: lesson.TeacherNotes,
		Homework:     lesson.Homework,
	}
}
