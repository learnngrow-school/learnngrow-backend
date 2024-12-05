package courses

type Course struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       int32 `json:"price" binding:"required"`
	Year        int16 `json:"grade" binding:"required"`
	Slug        string `json:"string"`
}

type CourseCreate struct {
	Course
	CategoryId  int32 `json:"categoryId" binding:"required"`
	SubjectId   int32 `json:"subjectId" binding:"required"`
}

type CourseWithData struct {
	Course   Course `json:"course"`
	Category Category `json:"category"`
	Subject  Subject `json:"subject"`
}

type CoursesGet []CourseWithData
