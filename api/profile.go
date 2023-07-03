package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/notify"
	"github.com/thomasfady/xsstower/notify/types"
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

// @BasePath	/api/
// @Summary	Get notification parameters
// @Description	Get notification parameters
// @Produce		json
// @Router			/profile/notifications [get]
// @Tags Profile
func GetNotificationParams(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user := models.GetUserById(user_id)
	var config []types.NotifierInformation
	for _, v := range notify.NotifiersList() {
		infos := v.GetBaseInformations()
		for k, p := range infos.Config {
			if !p.Sensitive {
				infos.Config[k].Value = user.NotifiersConfig[infos.Key+"."+p.Key]

			}

		}
		config = append(config, infos)

	}
	c.JSON(200, config)
}

// @BasePath	/api/
// @Summary	Set notification parameters
// @Description	Set notification parameters
// @Produce		json
// @Router			/profile/notifications [post]
// @Param			Params body []types.NotifierInformation	true	"Params"
// @Tags Profile
func PostNotificationParams(c *gin.Context) {
	var form []types.NotifierInformation
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		user_id := c.GetInt("user_id")
		user := models.GetUserById(user_id)

		if user.NotifiersConfig == nil {
			user.NotifiersConfig = make(map[string]string)
		}

		for _, notifier := range form {
			for _, param := range notifier.Config {
				if param.Value != "" {
					user.NotifiersConfig[notifier.Key+"."+param.Key] = param.Value
				}

			}
		}
		models.DB.Save(&user)
		var config []types.NotifierInformation
		for _, v := range notify.NotifiersList() {
			infos := v.GetBaseInformations()
			config = append(config, infos)
			for _, p := range infos.Config {
				if !p.Sensitive {
					p.Value = user.NotifiersConfig[infos.Key+"."+p.Key]
				}
			}

		}
		c.JSON(200, config)
	}

}
