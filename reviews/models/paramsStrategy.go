package reviews

import "learn-n-grow.dev/m/db/repository"

type CustomParams interface {
	GetType() string
	// GetBase() interface{}
}

func GetParams[T any](params CustomParams) T {
	if params.GetType() == "school" {
		return CreateSchoolReviewParams {}
	}
}
