package main

import (
	"embed"

	"log"

	"github.com/thomasfady/xsstower/config"
	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/routers"
	"github.com/thomasfady/xsstower/routes/unauthenticated"
	"github.com/thomasfady/xsstower/utils"

	"github.com/gin-gonic/gin"
)

//go:embed app
var assetsFS embed.FS

func main() {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	models.ConnectDatabase()

	config.ConfigPath = "app.yaml"
	config.LoadConfig()

	routers.SetupCors(r)
	routers.SetupSession(r)

	unauthenticated.AssetsFS = assetsFS
	utils.SeedAdmin()

	routers.UnauthenticatedRouter(r, "/api")
	routers.AuthenticatedRouter(r, "/api")
	routers.AdminRouter(r, "/api/admin")
	log.Printf("XssTower started on :8080")
	r.Run(":8080")
}
