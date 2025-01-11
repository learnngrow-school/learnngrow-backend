package lessons

type LessonCreate struct {
	StudentSlug  string `json:"studentSlug"`
	TeacherSlug  string `json:"teacherSlug"`
	Timestamp    int64  `json:"timestamp"`
	TeacherNotes string `json:"teacherNotes"`
	Homework     string `json:"homework"`
	Duration     int16  `json:"duration"`
	FileSlug     string `json:"fileSlug"`
}

type LessonGet struct {
	Timestamp    int64  `json:"timestamp"`
	TeacherNotes string `json:"teacherNotes"`
	Homework     string `json:"homework"`
	Duration     int16  `json:"duration"`
	FileSlug     string `json:"fileSlug"`
}
