package api

import (
	"net/http"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

// @BasePath	/api/
// @Summary	Handlers
// @Description	Get Handlers list
// @Produce		json
// @Router			/handlers [get]
// @Tags Handlers
func GetHandlers(c *gin.Context) {

	user_id := c.GetInt("user_id")

	handlers := models.GetViewableHandlers(user_id)
	c.JSON(http.StatusOK, handlers)
}

// @BasePath	/api/
// @Summary	Handlers
// @Description	Get Handler information
// @Produce		json
// @Router			/handlers/{handler_id} [get]
// @Tags Handlers
func GetHandler(c *gin.Context) {

	user_id := c.GetInt("user_id")
	handler := models.GetHandlerById(c.Param("key"))
	if handler.CanView(user_id) {
		c.JSON(http.StatusOK, handler)
	} else {
		c.Status(http.StatusForbidden)
	}

}

type HandlerUpdateForm struct {
	Screenshot    *bool    `json:"Screenshot"  binding:"required"`
	Dom           *bool    `json:"Dom"  binding:"required"`
	CollectedPage []string `json:"CollectedPages"`
}

type HandlerCreateForm struct {
	Screenshot *bool  `json:"screenshot"  binding:"required"`
	Dom        *bool  `json:"dom"  binding:"required"`
	Domain     string `json:"domain"  binding:"required"`
	Path       string `json:"path"  binding:"required"`
}

// @BasePath	/api/
// @Summary	Handlers
// @Description	Update Handler information
// @Produce		json
// @Router			/handlers/{handler_id} [put]
// @Param			HandlerUpdateForm body HandlerUpdateForm	true "HandleOptions"
// @Tags Handlers
func PutHandler(c *gin.Context) {
	var form HandlerUpdateForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		user_id := c.GetInt("user_id")

		handler := models.GetHandlerById(c.Param("key"))
		if handler.CanUpdate(user_id) {
			handler.Screenshot = *form.Screenshot
			handler.Dom = *form.Dom
			handler.CollectedPages = form.CollectedPage
			handler.Save()
			c.JSON(http.StatusOK, handler)
		} else {
			c.Status(http.StatusForbidden)
		}

	}

}

// @BasePath	/api/
// @Summary	Handlers
// @Description	Delete Handler
// @Produce		json
// @Router			/handlers/{handler_id} [delete]
// @Tags Handlers
func DeleteHandler(c *gin.Context) {
	user_id := c.GetInt("user_id")

	handler := models.GetHandlerById(c.Param("key"))
	if handler.CanDelete(user_id) {
		handler.Delete()
		c.JSON(http.StatusOK, "")
	} else {
		c.Status(http.StatusForbidden)
	}

}

// @BasePath	/api/
// @Summary	Create Handler
// @Description	Update Handler information
// @Produce		json
// @Router			/handlers [post]
// @Param			HandlerCreateForm body HandlerCreateForm	true "HandleOptions"
// @Tags Handlers
func PostHandler(c *gin.Context) {
	var form HandlerCreateForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	} else {
		user_id := c.GetInt("user_id")

		free, msg := models.HandleIsFree(form.Path, form.Domain)

		if !free {
			c.JSON(http.StatusForbidden, msg)
		} else {
			rbac := models.HandlerRbac{
				UserID:     user_id,
				Permission: models.OWNER,
			}
			new_handle := models.Handler{
				Domain:         form.Domain,
				Path:           form.Path,
				OwnerID:        user_id,
				Screenshot:     *form.Screenshot,
				Dom:            *form.Dom,
				CollectedPages: []string{},
				Members:        []models.HandlerRbac{rbac},
			}
			models.DB.Create(&new_handle)
			c.JSON(http.StatusOK, new_handle)
		}

	}

}
