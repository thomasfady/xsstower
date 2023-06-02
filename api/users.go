package api

import (
	"net/http"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

// @BasePath	/api/
// @Summary	Users
// @Description	Get Users list
// @Produce		json
// @Router			/users [get]
func GetUsers(c *gin.Context) {

	users := models.GetUsersByEmail(c.Query("search"))
	c.JSON(http.StatusOK, users)
}
