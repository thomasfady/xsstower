package tests

import (
	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/routers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func BaseRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	models.ConnectDatabase()

	store := cookie.NewStore([]byte("debug"))
	store.Options(sessions.Options{HttpOnly: true})
	router.Use(sessions.Sessions("session", store))

	routers.UnauthenticatedRouter(router, "/api")
	routers.AuthenticatedRouter(router, "/api")
	routers.AdminRouter(router, "/api/admin")
	return router
}
