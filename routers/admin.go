package routers

import (
	"github.com/thomasfady/xsstower/api/admin"
	"github.com/thomasfady/xsstower/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouter(engine *gin.Engine, path string) {
	group := engine.Group(path)

	group.Use(middlewares.AuthMiddleware())
	group.Use(middlewares.IsAdminMiddleware())
	group.GET("/users", admin.GetUsers)
	group.DELETE("/users/:key", admin.DeleteUser)
	group.PUT("/users/:key", admin.PutUser)
	group.POST("/users", admin.PostUsers)
	group.GET("/config", admin.GetConfig)
	group.PUT("/config", admin.PutConfig)
}
