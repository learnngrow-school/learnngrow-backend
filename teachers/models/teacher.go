package teachers

import auth "learn-n-grow.dev/m/auth/models"

type Teacher struct {
	SubjectIds []int32 `json:"subjectIds"`
	Biography  string  `json:"biography"`
}

type TeacherGet struct {
	User    auth.UserGet `json:"userData"`
	Teacher Teacher      `json:"teacherData"`
}

type TeachersGet []TeacherGet
