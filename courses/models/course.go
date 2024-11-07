package courses

type CourseCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" copier:"-"`
	Price       int32 `json:"price" binding:"required"`
	Year        int16 `json:"grade" binding:"required" copier:"-"`
	CategoryId  int32 `json:"categoryId" binding:"required"`
	SubjectId   int32 `json:"subjectId" binding:"required"`
}

type CourseGet struct {
	
}
