package unauthenticated

import (
	"embed"
	"mime"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var AssetsFS embed.FS

func GetApp(c *gin.Context) {
	index, _ := AssetsFS.ReadFile("app/index.html")
	file, err := AssetsFS.ReadFile("app" + c.Param("path"))
	if err != nil {
		c.Render(
			http.StatusOK, render.Data{
				ContentType: "text/html",
				Data:        index,
			})
	} else {
		c.Render(
			http.StatusOK, render.Data{
				ContentType: mime.TypeByExtension(filepath.Ext(c.Param("path"))),
				Data:        file,
			})
	}

}
