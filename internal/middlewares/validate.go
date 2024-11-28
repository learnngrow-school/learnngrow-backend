package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	// Validator "github.com/go-playground/validator/v10"
	"learn-n-grow.dev/m/utils"
)

type Validatable interface {
	Validate() error
}

func Validate[T Validatable](model T) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var input T

		if err = c.ShouldBindJSON(&input); err != nil {
			utils.Throw(c,
				http.StatusUnprocessableEntity,
				errors.New("Bad request data"),
			)
		}

		c.Set("input", input)

		if err = input.Validate(); err != nil {
			utils.Throw(c,
				http.StatusUnprocessableEntity,
				err,
			)
		}
	}
}
