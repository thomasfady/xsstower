package routers

import (
	"github.com/thomasfady/xsstower/api"
	"github.com/thomasfady/xsstower/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthenticatedRouter(engine *gin.Engine, path string) {
	group := engine.Group(path)

	group.Use(middlewares.AuthMiddleware())

	group.GET("/users", api.GetUsers)
	group.GET("/payloads", api.GetPayloads)
	group.GET("/fires", api.GetFires)
	group.GET("/fire/:key", api.GetFire)
	group.GET("/fire/:key/file/:id", api.GetFireFile)
	group.POST("/fire/:key/share", api.PostFireShare)
	group.DELETE("/fire/:key/share", api.DeleteFireShare)
	group.DELETE("/fire/:key", api.DeleteFire)
	group.GET("/handlers", api.GetHandlers)
	group.POST("/handlers", api.PostHandler)
	group.POST("/handler/:key/members", api.PostHandlerMembers)
	group.PUT("/handler/:key/members", api.PutHandlerMembers)
	group.DELETE("/handler/:key/members", api.DeleteHandlerMembers)

	group.GET("/handler/:key", api.GetHandler)
	group.PUT("/handler/:key", api.PutHandler)
	group.DELETE("/handler/:key", api.DeleteHandler)

	group.POST("/profile/set_password", api.PostSetPassword)
	group.GET("/profile/notifications", api.GetNotificationParams)
	group.POST("/profile/notifications", api.PostNotificationParams)
}
