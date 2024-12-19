package reviews

import validation "github.com/go-ozzo/ozzo-validation"

type ReviewCreate struct {
	Details    string `json:"details"`
	AuthorName string `json:"authorName"`
}

type ReviewGet struct {
	ID int `json:"id"`
	ReviewCreate
}

type TeacherReviewCreate struct {
	Rating     int16  `json:"rating"`
	Details    string `json:"details"`
	AuthorName string `json:"authorName"`
}

type TeacherReviewGet struct {
	ID int `json:"id"`
	ReviewCreate
}

type CourseReviewCreate struct {
	Rating     int16  `json:"rating"`
	Details    string `json:"details"`
	AuthorName string `json:"authorName"`
}

type CourseReviewGet struct {
	ID int `json:"id"`
	ReviewCreate
}


func (review ReviewCreate) Validate() error {
	return validation.ValidateStruct(&review,
		validation.Field(&review.Details, validation.Required, validation.Length(8, 1023)),
		validation.Field(&review.AuthorName, validation.Required, validation.Length(3, 255)),
	)
}


func (review TeacherReviewCreate) Validate() error {
	return validation.ValidateStruct(&review,
		validation.Field(&review.Details, validation.Required, validation.Length(8, 1023)),
		validation.Field(&review.AuthorName, validation.Required, validation.Length(3, 255)),
	)
}
