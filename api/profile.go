package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/notify"
	"github.com/thomasfady/xsstower/notify/types"
)

type SetPasswordForm struct {
	CurrentPassword string `json:"current_password"  binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// @BasePath		/api/
// @Summary		Change Password
// @Description	Change Password
// @Produce		json
// @Router			/profile/set_password [post]
// @Param			SetPasswordForm	body	SetPasswordForm	true	"Credentials"
// @Tags			Profile
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

type SetSecondFactorForm struct {
	OTP    string `json:"otp_code"  binding:"required"`
	Secret string `json:"secret" binding:"required"`
	Enable bool   `json:"enable" binding:"required"`
}

// @BasePath		/api/
// @Summary		Set 2FA
// @Description	Change Password
// @Produce		json
// @Router			/profile/2fa [post]
// @Param			SetSecondFactorForm	body	SetSecondFactorForm	true	"2FA config"
// @Tags			Profile
func PostSetSecondFactor(c *gin.Context) {
	var form SetSecondFactorForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		user_id := c.GetInt("user_id")
		user := models.GetUserById(user_id)
		if !totp.Validate(form.OTP, form.Secret) {
			c.JSON(http.StatusForbidden, "OTP not valid")
		} else {
			user.OTPEnabled = true
			user.OTPSecret = form.Secret
			models.DB.Save(&user)
			c.Status(200)
		}
	}
}

// @BasePath		/api/
// @Summary		Delete 2FA
// @Description	Delete 2FA
// @Produce		json
// @Router			/profile/2fa [delete]
// @Tags			Profile
func DeleteSecondFactor(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user := models.GetUserById(user_id)

	user.OTPEnabled = false
	user.OTPSecret = ""
	models.DB.Save(&user)
	c.Status(200)
}

type GetSecondFactorResp struct {
	Enabled bool   `json:"enabled"`
	URI     string `json:"totp_uri,omitempty"`
	Secret  string `json:"secret,omitempty"`
}

// @BasePath		/api/
// @Summary		Get 2FA information
// @Description	Get 2FA information
// @Produce		json
// @Router			/profile/2fa [get]
// @Tags			Profile
func GetSecondFactor(c *gin.Context) {
	user_id := c.GetInt("user_id")
	user := models.GetUserById(user_id)
	resp := GetSecondFactorResp{
		Enabled: user.OTPEnabled,
	}
	if !user.OTPEnabled {
		secret, _ := totp.Generate(totp.GenerateOpts{
			Issuer:      "XSSTower",
			AccountName: user.Email,
		})
		resp.Secret = secret.Secret()
		resp.URI = secret.URL()
	}
	c.JSON(200, resp)
}

// @BasePath		/api/
// @Summary		Get notification parameters
// @Description	Get notification parameters
// @Produce		json
// @Router			/profile/notifications [get]
// @Tags			Profile
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

// @BasePath		/api/
// @Summary		Set notification parameters
// @Description	Set notification parameters
// @Produce		json
// @Router			/profile/notifications [post]
// @Param			Params	body	[]types.NotifierInformation	true	"Params"
// @Tags			Profile
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
