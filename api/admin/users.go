package admin

import (
	"net/http"
	"strconv"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

// @BasePath	/api/
// @Summary	Users
// @Description	Get Users list
// @Produce		json
// @Router			/admin/users [get]
// @Tags Administration
func GetUsers(c *gin.Context) {

	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @BasePath	/api/
// @Summary	Users
// @Description	Delete User
// @Produce		json
// @Router			/admin/users/{user_id} [delete]
// @Tags Administration
func DeleteUser(c *gin.Context) {
	user_id := c.GetInt("user_id")

	if c.Param("key") == strconv.Itoa(user_id) {
		c.JSON(http.StatusUnauthorized, "Unable to delete your own user")
	} else {
		var user models.User
		models.DB.First(&user, " id = ?", c.Param("key"))
		models.DB.Delete(&user)
		c.JSON(http.StatusOK, "")
	}
}

type UserUpdateForm struct {
	IsAdmin  *bool
	Password string
}

// @BasePath	/api/
// @Summary	Users
// @Description	Update User information
// @Produce		json
// @Router			/admin/users/{user_id} [put]
// @Param			UserUpdateForm body UserUpdateForm	true "UserOptions"
// @Tags Administration
func PutUser(c *gin.Context) {
	var form UserUpdateForm

	user_id := c.GetInt("user_id")

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		if c.Param("key") == strconv.Itoa(user_id) {
			c.JSON(http.StatusUnauthorized, "Unable to modify your own user")
		} else {
			var user models.User
			models.DB.First(&user, "id = ?", c.Param("key"))
			if user.ID == 0 {
				c.Status(http.StatusNotFound)
			} else {
				if form.IsAdmin != nil {
					user.IsAdmin = *form.IsAdmin
				}
				if form.Password != "" {
					user.SetPassword(form.Password)
				}
				models.DB.Save(user)
				c.JSON(http.StatusOK, "")
			}
		}

	}

}

type UserCreateForm struct {
	IsAdmin  *bool  `binding:"required"`
	Password string `binding:"required"`
	Email    string `binding:"required"`
	Username string `binding:"required"`
}

// @BasePath	/api/
// @Summary	Create User
// @Description	Create User
// @Produce		json
// @Router			/admin/users [post]
// @Param			UserCreateForm body UserCreateForm	true "UserOptions"
// @Tags Administration
func PostUsers(c *gin.Context) {
	var form UserCreateForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		var user models.User
		models.DB.First(&user, "username = ?", form.Username)
		if user.ID != 0 {
			c.JSON(http.StatusNotAcceptable, "Username already taken")
			return
		}

		models.DB.First(&user, "email = ?", form.Email)
		if user.ID != 0 {
			c.JSON(http.StatusNotAcceptable, "Email already taken")
			return
		}

		new_user := models.User{
			Email:    form.Email,
			Username: form.Username,
			IsAdmin:  *form.IsAdmin,
		}
		new_user.SetPassword(form.Password)
		models.DB.Create(&new_user)
		c.JSON(http.StatusOK, new_user)
	}

}
