package api

import (
	"net/http"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @BasePath	/api/
// @Summary	Login
// @Description	Login
// @Produce		json
// @Router			/login [post]
// @Param			LoginForm body LoginForm	true	"Credentials"
// @Tags Authentication
func PostLogin(c *gin.Context) {
	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {

		user := models.GetUserByEmail(form.Email)

		if user.CheckPassword(form.Password) {
			session := sessions.Default(c)

			session.Set("user_id", user.ID)
			session.Set("is_admin", user.IsAdmin)
			session.Save()
			c.JSON(http.StatusOK, user)
			return

		} else {
			c.JSON(http.StatusUnauthorized, "Bad email or password")
		}
	}
}
