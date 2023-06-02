package routers

import (
	"github.com/thomasfady/xsstower/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupSession(engine *gin.Engine) {
	var store cookie.Store
	if gin.Mode() == "debug" {
		store = cookie.NewStore([]byte("debug"))
	} else {
		store = cookie.NewStore([]byte(utils.RandomHex(20)))
	}
	store.Options(sessions.Options{HttpOnly: true})
	engine.Use(sessions.Sessions("session", store))
}

func SetupCors(engine *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Content-Length", "Content-Type", "XssTower-File", "XssTower-Key"}
	engine.Use(cors.New(config))
}
