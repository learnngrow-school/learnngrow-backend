package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	Validator "github.com/go-playground/validator/v10"
)

// TODO: implement
func ValidationMiddleware(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error

		if err = c.ShouldBindJSON(&model); err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity,
				gin.H{},
			)
		}

		validator := Validator.New()
		if err = validator.Struct(model); err != nil {
			validationErrors := err.(Validator.ValidationErrors)
			errorStrings := make([]string, len(validationErrors))
			c.AbortWithStatusJSON(
				http.StatusUnprocessableEntity, errorStrings,
			)
		}
	}
}
