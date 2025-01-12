package lessons

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jinzhu/copier"
	"learn-n-grow.dev/m/db/repository"
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

	var lesson models.LessonCreate
	if err = c.ShouldBindJSON(&lesson); err != nil {
		utils.Throw(c, http.StatusBadRequest, err)
		return
	}

	params := GetCreateLessonParams(lesson)
	result, err := utils.GetRepo(c).CreateLesson(c, params)
	if err != nil {
		utils.Throw(c, http.StatusInternalServerError, err)
		return
	}

	var lessonResult models.LessonGet
	copier.CopyWithOption(&lessonResult, &result, copier.Option{DeepCopy: true})
	lessonResult.Timestamp = lesson.Timestamp

	c.JSON(http.StatusCreated, lessonResult)
}

func GetCreateLessonParams(lesson models.LessonCreate) repository.CreateLessonParams {
	slug := pgtype.Text{
		String: lesson.FileSlug,
		Valid:  len(lesson.FileSlug) > 5,
	}
	return repository.CreateLessonParams{
		Slug:         lesson.StudentSlug,
		Slug_2:       lesson.TeacherSlug,
		Ts:           pgtype.Timestamptz{Time: time.Unix(lesson.Timestamp, 0), Valid: true},
		TeacherNotes: lesson.TeacherNotes,
		Homework:     lesson.Homework,
		Duration:     lesson.Duration,
		FileSlug:     slug,
	}
}
