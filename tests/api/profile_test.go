package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomasfady/xsstower/api"
	"github.com/thomasfady/xsstower/models"
	"github.com/thomasfady/xsstower/tests"
)

func TestPostSetPassword(t *testing.T) {
	tests.ResetDB()
	router := tests.BaseRouter()
	user := models.User{
		Email:    "user@user.fr",
		Username: "user",
	}
	user.SetPassword("user")
	models.DB.Create(&user)

	requester := tests.NewRequester(router)

	body := api.LoginForm{
		Email:    "user@user.fr",
		Password: "user",
	}
	res := requester.Post("/api/login", body)
	assert.Equal(t, 200, res.Code)

	body_set_pass := api.SetPasswordForm{
		CurrentPassword: "dummy",
		NewPassword:     "dummy1",
		ConfirmPassword: "dummy2",
	}
	res = requester.Post("/api/profile/set_password", body_set_pass)
	assert.Equal(t, 403, res.Code)

	body_set_pass.CurrentPassword = "user"

	res = requester.Post("/api/profile/set_password", body_set_pass)
	assert.Equal(t, 406, res.Code)

	body_set_pass.ConfirmPassword = "dummy1"

	res = requester.Post("/api/profile/set_password", body_set_pass)
	assert.Equal(t, 200, res.Code)

	requester = tests.NewRequester(router)
	res = requester.Post("/api/login", body)
	assert.Equal(t, 401, res.Code)

	body.Password = "dummy1"
	res = requester.Post("/api/login", body)
	assert.Equal(t, 200, res.Code)
}
