package api

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

// @BasePath	/api/
// @Summary	Public Fire
// @Description	Get XSS Fire information
// @Produce		json
// @Router			/sharing/{fire_id} [get]
func GetPublicFire(c *gin.Context) {

	hit := models.GetHitBySharingKey(c.Param("key"))
	if hit.ID != 0 {
		c.JSON(http.StatusOK, hit)
	} else {
		c.Status(404)
	}

}

// @BasePath	/api/
// @Summary	Public Fire
// @Description	Get XSS Fire File
// @Router			/sharing/{fire_id}/file/{file_id} [get]
func GetPublicFireFile(c *gin.Context) {

	hit := models.GetHitBySharingKey(c.Param("key"))
	if hit.ID != 0 {
		for _, cp := range hit.CollectedPages {
			if strconv.FormatInt(int64(cp.ID), 10) == c.Param("id") {
				c.Header("Content-Disposition", "attachment; filename=\""+filepath.Base(cp.Path)+"\"")
				c.Data(http.StatusOK, "application/data", []byte(cp.Content))
				return
			}
		}
		c.Status(404)
	} else {
		c.Status(404)
	}

}
