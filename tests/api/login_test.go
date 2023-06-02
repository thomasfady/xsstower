package api

import (
	"testing"

	"github.com/thomasfady/xsstower/api"
	"github.com/thomasfady/xsstower/middlewares"
	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := tests.BaseRouter()

	api_authenticated := router.Group("/api")
	api_authenticated.Use(middlewares.AuthMiddleware())
	api_authenticated.GET("/mock_authenticated", func(c *gin.Context) {
		c.Status(200)
	})

	api_admin := router.Group("/api/admin")
	api_admin.Use(middlewares.AuthMiddleware())
	api_admin.Use(middlewares.IsAdminMiddleware())
	api_admin.GET("/mock_admin", func(c *gin.Context) {
		c.Status(200)
	})
	return router
}

func TestPostLogin(t *testing.T) {
	router := SetUpRouter()

	requester := tests.NewRequester(router)

	res := requester.Post("/api/login", nil)

	assert.Equal(t, 406, res.Code)

	body := api.LoginForm{
		Email:    "user@user.fr",
		Password: "user",
	}

	res = requester.Post("/api/login", body)
	assert.Equal(t, 401, res.Code)
	assert.Equal(t, "Bad email or password", res.ResponseString())

	user := models.User{
		Email:    "user@user.fr",
		Username: "user",
	}
	user.SetPassword("test")
	models.DB.Create(&user)

	res = requester.Post("/api/login", body)
	assert.Equal(t, 401, res.Code)
	assert.Equal(t, "Bad email or password", res.ResponseString())

	body.Password = "test"

	res = requester.Post("/api/login", body)
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "{\"Email\":\"user@user.fr\",\"Username\":\"user\",\"IsAdmin\":false}", res.ResponseString())

	res = requester.Get("/api/mock_authenticated")
	assert.Equal(t, 200, res.Code)

	res = requester.Get("/api/admin/mock_admin")
	assert.Equal(t, 403, res.Code)

	user.IsAdmin = true
	models.DB.Save(user)

	requester.Post("/api/login", body)

	res = requester.Get("/api/admin/mock_admin")
	assert.Equal(t, 200, res.Code)
}
