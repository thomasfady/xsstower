package tests

import "github.com/thomasfady/xsstower/models"

func InsertUser(email string, password string) (user models.User) {
	user = models.User{
		Email:    email,
		Username: email,
	}
	user.SetPassword(password)
	models.DB.Create(&user)
	return
}

func ResetDB() {
	models.DB.Exec("DELETE FROM users")
}
