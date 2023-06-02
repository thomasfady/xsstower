package api

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

type APIFires struct {
	Ip             string
	Handler        string
	Url            string
	UserAgent      string
	Referer        string
	CorrelationKey string
}

// @BasePath	/api/
// @Summary	Fires
// @Description	Get XSS Fires
// @Produce		json
// @Tags XSSFire
// @Router			/fires [get]
func GetFires(c *gin.Context) {

	user_id := c.GetInt("user_id")

	fires := make([]APIFires, 0)

	var pagination models.Pagination
	c.Bind(&pagination)

	hits := models.GetViewableHitPaginated(user_id, &pagination)

	for _, v := range hits {
		fires = append(fires, APIFires{
			Ip:             v.Ip,
			Handler:        v.Handler.Path,
			Url:            v.Url,
			UserAgent:      v.UserAgent,
			Referer:        v.Referer,
			CorrelationKey: v.CorrelationKey,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"fires":       fires,
		"limit":       pagination.Limit,
		"page":        pagination.Page,
		"total_rows":  pagination.TotalRows,
		"total_pages": pagination.TotalPages,
	})
}

// @BasePath	/api/
// @Summary	Fire
// @Description	Get XSS Fire information
// @Produce		json
// @Tags XSSFire
// @Router			/fire/{fire_id} [get]
func GetFire(c *gin.Context) {

	user_id := c.GetInt("user_id")

	hit := models.GetHitByCorrelationKey(c.Param("key"))
	if hit.CanView(user_id) {
		c.JSON(http.StatusOK, hit)
	} else {
		c.JSON(http.StatusForbidden, nil)
	}

}

// @BasePath	/api/
// @Summary	Fire
// @Tags XSSFire
// @Description	Get XSS Fire File
// @Router			/fire/{fire_id}/file/{file_id} [get]
func GetFireFile(c *gin.Context) {

	user_id := c.GetInt("user_id")

	hit := models.GetHitByCorrelationKey(c.Param("key"))
	if hit.ID == 0 {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if hit.CanView(user_id) {
		for _, cp := range hit.CollectedPages {
			if strconv.FormatInt(int64(cp.ID), 10) == c.Param("id") {
				c.Header("Content-Disposition", "attachment; filename=\""+filepath.Base(cp.Path)+"\"")
				c.Data(http.StatusOK, "application/data", []byte(cp.Content))
				return
			}
		}
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusForbidden, nil)
	}
}

// @BasePath	/api/
// @Summary	Fires
// @Description	Delete Fire
// @Produce		json
// @Tags XSSFire
// @Router			/fire/{fire_id} [delete]
func DeleteFire(c *gin.Context) {
	user_id := c.GetInt("user_id")

	hit := models.GetHitByCorrelationKey(c.Param("key"))
	if hit.CanDelete(user_id) {
		hit.Delete()
	} else {
		c.JSON(http.StatusForbidden, nil)
	}

	c.JSON(http.StatusOK, "")

}

// @BasePath	/api/
// @Summary	Fires sharing
// @Description	Enable fire sharing
// @Produce		json
// @Tags XSSFire
// @Router			/fire/{fire_id}/share [post]
func PostFireShare(c *gin.Context) {

	user_id := c.GetInt("user_id")

	hit := models.GetHitByCorrelationKey(c.Param("key"))
	if hit.CanUpdate(user_id) {
		hit.EnableSharing()
		c.Status(http.StatusCreated)
	} else {
		c.Status(http.StatusForbidden)
	}
}

// @BasePath	/api/
// @Summary	Fires sharing
// @Description	Disable fire sharing
// @Tags XSSFire
// @Router			/fire/{fire_id}/share [delete]
func DeleteFireShare(c *gin.Context) {

	user_id := c.GetInt("user_id")

	hit := models.GetHitByCorrelationKey(c.Param("key"))
	if hit.CanUpdate(user_id) {
		hit.DisableSharing()
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusForbidden)
	}
}
