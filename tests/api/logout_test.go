package api

import (
	"testing"

	"github.com/thomasfady/xsstower/api"
	"github.com/thomasfady/xsstower/middlewares"
	"github.com/thomasfady/xsstower/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouterForLogout() *gin.Engine {
	tests.ResetDB()
	router := tests.BaseRouter()

	api_authenticated := router.Group("/api")
	api_authenticated.Use(middlewares.AuthMiddleware())
	api_authenticated.GET("/mock_authenticated", func(c *gin.Context) {
		c.Status(200)
	})
	return router
}

func TestPostLogout(t *testing.T) {
	router := SetUpRouterForLogout()

	requester := tests.NewRequester(router)

	tests.InsertUser("user@user.fr", "user")

	body := api.LoginForm{
		Email:    "user@user.fr",
		Password: "user",
	}

	res := requester.Post("/api/login", body)
	assert.Equal(t, 200, res.Code)

	res = requester.Get("/api/mock_authenticated")
	assert.Equal(t, 200, res.Code)

	res = requester.Post("/api/logout", body)
	assert.Equal(t, 200, res.Code)

	res = requester.Get("/api/mock_authenticated")
	assert.Equal(t, 401, res.Code)

}
