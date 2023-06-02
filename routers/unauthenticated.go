package routers

import (
	"github.com/thomasfady/xsstower/api"
	"github.com/thomasfady/xsstower/docs"
	"github.com/thomasfady/xsstower/routes/unauthenticated"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func UnauthenticatedRouter(engine *gin.Engine, path string) {
	docs.SwaggerInfo.BasePath = "/api"

	engine.GET("/:slug", unauthenticated.GetPayload)
	engine.POST("/:slug", unauthenticated.PostPayload)
	engine.PUT("/:slug", unauthenticated.PutPayload)
	engine.GET("/", unauthenticated.GetRoot)
	engine.GET("/app/*path", unauthenticated.GetApp)

	group := engine.Group(path)
	group.POST("/login", api.PostLogin)
	group.POST("/logout", api.PostLogout)
	group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	group.GET("/sharing/:key", api.GetPublicFire)
	group.GET("/sharing/:key/file/:id", api.GetPublicFireFile)
}
