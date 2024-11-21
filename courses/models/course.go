package courses

type CourseCreate struct {
	Course
	// Title       string `json:"title" binding:"required"`
	// Description string `json:"description" copier:"-"`
	// Price       int32 `json:"price" binding:"required"`
	// Year        int16 `json:"grade" binding:"required" copier:"-"`
	CategoryId  int32 `json:"categoryId" binding:"required"`
	SubjectId   int32 `json:"subjectId" binding:"required"`
}

type Course struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price       int32 `json:"price" binding:"required"`
	Year        int16 `json:"grade" binding:"required"`
}

type CourseWithData struct {
	Course   Course `json:"course"`
	Category Category `json:"category"`
	Subject  Subject `json:"subject"`
}

type CoursesGet []CourseWithData
