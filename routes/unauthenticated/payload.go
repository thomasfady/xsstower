package unauthenticated

import (
	"embed"
	"encoding/json"
	"html/template"
	"io"
	"strings"

	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//go:embed templates
var templatesFS embed.FS

var payloadTemplate *template.Template

func initPayloadTemplate() {

}

type PayloadData struct {
	CorrelationKey string
	ServerUrl      string
	Callback       string
	Screenshot     bool
	Dom            bool
	CollectPages   string
}

func GetPayload(c *gin.Context) {
	slug := c.Param("slug")
	hostname := strings.Split(c.Request.Host, ":")[0]
	var handler models.Handler
	models.DB.Where("path = ? and domain = ?", slug, hostname).Or("path = ? and domain = ?", slug, "*").Or("path = ? and domain = ?", "*", hostname).First(&handler)

	if handler.ID == 0 {
		c.String(404, "Not Found")
		c.Abort()
		return
	}
	id := uuid.New()
	hit := &models.XssHit{
		CorrelationKey: id.String(),
		Ip:             c.ClientIP(),
		UserAgent:      c.Request.Header.Get("User-Agent"),
		Origin:         c.Request.Header.Get("Origin"),
		Referer:        c.Request.Header.Get("Referer"),
		OwnerID:        handler.OwnerID,
		HandlerID:      int(handler.ID),
	}
	models.DB.Create(hit)

	type Data struct {
		CorrelationKey string
		ServerUrl      string
		Callback       string
		Screenshot     bool
		Dom            bool
		CollectPages   string
	}
	cp, _ := json.Marshal(handler.CollectedPages)
	ctx := &PayloadData{
		CorrelationKey: hit.CorrelationKey,
		ServerUrl:      utils.BaseUrl(c),
		Callback:       "/" + slug,
		Screenshot:     handler.Screenshot,
		Dom:            handler.Dom,
		CollectPages:   string(cp),
	}

	tmpl, err := template.New("payload.js").Funcs(template.FuncMap{
		"unescapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"include": func(s string) string {
			file, err := templatesFS.ReadFile("templates/" + s)
			if err != nil {
				panic(err)
			}
			return string(file)
		},
	}).ParseFS(templatesFS, "templates/payload.js")
	if err != nil {
		panic(err)
	}
	c.Header("Content-Type", "text/js")
	tmpl.Execute(c.Writer, ctx)
}

type CallbackModel struct {
	Cookies        string `json:"cookies"`
	CorrelationKey string `json:"key"`
	Local          string `json:"local"`
	Session        string `json:"session"`
	Url            string `json:"url"`
	Dom            string `json:"dom"`
	Screenshot     string `json:"screenshot"`
}

func PostPayload(c *gin.Context) {
	var params CallbackModel
	c.Bind(&params)
	var hit models.XssHit
	models.DB.First(&hit, "correlation_key = ?", params.CorrelationKey)

	if params.Dom != "" {
		hit.Dom = params.Dom
	} else if params.Screenshot != "" {
		hit.Screenshot = params.Screenshot
	} else {
		hit.Cookies = params.Cookies
		hit.LocalStorage = params.Local
		hit.SessionStorage = params.Session
		hit.Url = params.Url
	}

	models.DB.Save(hit)
}

func PutPayload(c *gin.Context) {
	correlation_key := c.Request.Header["Xsstower-Key"]
	path := c.Request.Header["Xsstower-File"]
	var hit models.XssHit
	models.DB.First(&hit, "correlation_key = ?", correlation_key[0])
	if hit.ID != 0 {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(400)
		} else {
			cp := &models.CollectedPage{
				XssHitID: hit.ID,
				Path:     path[0],
				Content:  body,
			}
			models.DB.Create(&cp)
			c.Status(200)
			c.Abort()
		}

	} else {
		c.Status(404)
	}
}
