package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @BasePath	/api/
// @Summary	Config
// @Description	Get Config
// @Produce		json
// @Router			/admin/config [get]
// @Tags Administration
func GetConfig(c *gin.Context) {

	c.JSON(http.StatusOK, viper.AllSettings())
}

// @BasePath	/api/
// @Summary	Config
// @Description	Update config Config
// @Produce		json
// @Router			/admin/config [put]
// @Param			ConfigValues body map[string]string true "ConfigValues"
// @Success      200  {object}  map[string]interface{}
// @Tags Administration
func PutConfig(c *gin.Context) {
	var obj map[string]interface{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	json.Unmarshal(body, &obj)
	for k, v := range obj {
		if viper.IsSet(k) {
			viper.Set(k, v)
		}
	}
	viper.WriteConfig()
	c.JSON(http.StatusOK, viper.AllSettings())
}
