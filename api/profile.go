package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomasfady/xsstower/models"
)

type SetPasswordForm struct {
	CurrentPassword string `json:"current_password"  binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// @BasePath	/api/
// @Summary	Change Password
// @Description	Change Password
// @Produce		json
// @Router			/profile/set_password [post]
// @Param			SetPasswordForm body SetPasswordForm	true	"Credentials"
// @Tags Authentication
func PostSetPassword(c *gin.Context) {
	var form SetPasswordForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		user_id := c.GetInt("user_id")
		user := models.GetUserById(user_id)
		if user.CheckPassword(form.CurrentPassword) {
			if form.NewPassword == form.ConfirmPassword {
				user.SetPassword(form.NewPassword)
				models.DB.Save(&user)
				c.Status(200)
			} else {
				c.JSON(http.StatusNotAcceptable, "Passwords mismatch")
			}
		} else {
			c.JSON(http.StatusForbidden, "Old password not valid")
		}

	}
}
