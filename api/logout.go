package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// @BasePath	/api/
// @Summary	Logout
// @Description	Logout Not working for now
// @Produce		json
// @Router			/logout [post]
func PostLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})

}
