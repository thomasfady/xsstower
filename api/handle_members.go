package api

import (
	"net/http"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

type HandlerMemberCreateForm struct {
	Role  models.PermisionType `json:"role"  binding:"required"`
	Email string               `json:"email"  binding:"required"`
}

type HandlerMemberEditForm struct {
	Role models.PermisionType `json:"role"  binding:"required"`
	Id   int                  `json:"user_id"  binding:"required"`
}
type HandlerMemberDeleteForm struct {
	Id int `json:"user_id"  binding:"required"`
}

// @BasePath	/api/
// @Summary	Update Handler Member Role
// @Description	Update Handler Member Role
// @Produce		json
// @Router			/handler/{handler_id}/members [put]
// @Tags Handlers
// @Param			HandlerMemberEditForm body HandlerMemberEditForm	true "HandleOptions"
func PutHandlerMembers(c *gin.Context) {
	var form HandlerMemberEditForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	user_id := c.GetInt("user_id")

	var handler models.Handler
	models.DB.Preload("Members.User").First(&handler, "owner_id = ? and id = ?", user_id, c.Param("key"))

	if handler.ID == 0 {
		c.JSON(http.StatusNotFound, "Handler does not exists")
		return
	}

	if user_id == form.Id {
		c.JSON(http.StatusForbidden, "You can't update your role")
		return
	}
	for _, member := range handler.Members {
		if member.UserID == form.Id {
			member.Permission = models.PermisionType(form.Role)
			models.DB.Save(&member)
			c.JSON(http.StatusOK, "Member role updatated")
			return
		}
	}
}

// @BasePath	/api/
// @Summary	Delete Handler Member
// @Description	Delete Handler Member
// @Produce		json
// @Router			/handler/{handler_id}/members [delete]
// @Tags Handlers
// @Param			HandlerMemberDeleteForm body HandlerMemberDeleteForm	true "HandleOptions"
func DeleteHandlerMembers(c *gin.Context) {
	var form HandlerMemberDeleteForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	user_id := c.GetInt("user_id")

	var handler models.Handler
	models.DB.Preload("Members.User").First(&handler, "owner_id = ? and id = ?", user_id, c.Param("key"))

	if handler.ID == 0 {
		c.JSON(http.StatusNotFound, "Handler does not exists")
		return
	}

	if user_id == form.Id {
		c.JSON(http.StatusForbidden, "You can't delete yourself")
		return
	}
	for _, member := range handler.Members {
		if member.UserID == form.Id {
			models.DB.Delete(&member)
			c.JSON(http.StatusOK, "Member deleted")
			return
		}
	}
}

// @BasePath	/api/
// @Summary	Create Handler Member
// @Description	Create Handler Member
// @Produce		json
// @Router			/handler/{handler_id}/members [post]
// @Tags Handlers
// @Param			HandlerMemberCreateForm body HandlerMemberCreateForm	true "HandleOptions"
func PostHandlerMembers(c *gin.Context) {
	var form HandlerMemberCreateForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	user_id := c.GetInt("user_id")

	var handler models.Handler
	models.DB.Preload("Members.User").First(&handler, "owner_id = ? and id = ?", user_id, c.Param("key"))

	var user models.User
	models.DB.First(&user, "email = ?", form.Email)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, "User does not exists")
		return
	}
	if user.ID == uint(user_id) {
		c.JSON(http.StatusNotAcceptable, "Member already exists")
		return
	}
	for _, member := range handler.Members {
		if user.ID == member.User.ID {
			c.JSON(http.StatusNotAcceptable, "Member already exists")
			return
		}
	}
	rbac := models.HandlerRbac{
		User:       user,
		Permission: form.Role,
	}
	handler.Members = append(handler.Members, rbac)
	tx := models.DB.Save(handler)
	if tx.Error != nil {
		c.JSON(http.StatusNotAcceptable, "Error during Handler Member creation")
	}

}
