package admin

import validation "github.com/go-ozzo/ozzo-validation"

type Subject struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func (subject Subject) Validate() error {
	return validation.ValidateStruct(&subject,
		validation.Field(&subject.Title, validation.Required, validation.Length(3, 63)),
	)
}
