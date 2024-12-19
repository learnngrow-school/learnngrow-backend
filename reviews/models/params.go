package reviews

import "learn-n-grow.dev/m/db/repository"


type CreateSchoolReviewParams struct {
	repository.CreateSchoolReviewParams
	// Details    string
	// AuthorName string
}

func (params *CreateSchoolReviewParams) GetType() string {
	return "CreateSchoolReviewParams"
}
