package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"learn-n-grow.dev/m/internal"
)

// Logout Logs user out (removes token cookie)
// @summary Logs user out (removes token cookie)
// @produce json
// @tags    Auth
// @success 200 {object} string
// @router  /auth/logout [post]
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", internal.Server.Domain, false, true)
	c.JSON(http.StatusOK, "ok")
}
