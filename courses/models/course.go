package courses

import validation "github.com/go-ozzo/ozzo-validation"

type CourseCreate struct {
	Course
	CategoryId  int32 `json:"categoryId" binding:"required"`
	SubjectId   int32 `json:"subjectId" binding:"required"`
}

type Course struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" copier:"-"`
	Price       int32 `json:"price" binding:"required"`
	Year        int16 `json:"grade" binding:"required" copier:"-"`
}

type CourseWithData struct {
	Course   Course `json:"course"`
	Category Category `json:"category"`
	Subject  Subject `json:"subject"`
}

type CoursesGet []CourseWithData


func (course CourseCreate) Validate() error {
	return validation.ValidateStruct(&course,
		validation.Field(&course.Title, validation.Required, validation.Length(8, 255)),
		validation.Field(&course.Description, validation.Required, validation.Length(8, 4095)),
		validation.Field(&course.Year, validation.Required, validation.Min(0), validation.Max(11)),
	)
}
