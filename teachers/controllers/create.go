package teachers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "learn-n-grow.dev/m/auth/controllers"
	"learn-n-grow.dev/m/utils"
)

// Create   Create teacher
// @summary Create teacher
// @accept  json
// @produce json
// @param   user body auth.UserCreate true "User"
// @tags    Teachers
// success  201 {object} auth.UserGet
// @router  /teachers/ [post]
func Create(c *gin.Context) {
	email, emailIsSet := c.Get("x-email")
	if !emailIsSet || email != "admin" {
		utils.Throw(c, http.StatusUnauthorized, errors.New("You are not superuser"))
		return
	}
	auth.CreateUser(c, true)
}
