package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Email        string        `gorm:"size:255;not null;unique"`
	Username     string        `gorm:"size:255;not null;unique"`
	Password     string        `gorm:"size:255;not null;" json:"-"`
	IsAdmin      bool          `gorm:"default:false"`
	HandlersRbac []HandlerRbac `json:"-"`
}

func (u *User) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)) == nil
}

func (u *User) SetPassword(pass string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func GetUserById(user_id int) User {
	var user User
	DB.First(&user, user_id)
	return user
}

func (u *User) HandlerRbacs() []HandlerRbac {

	rbacs := []HandlerRbac{}

	DB.Preload("Handler").Find(&rbacs, "user_id = ?", u.ID)

	return rbacs
}

func GetUserByEmail(email string) (user User) {
	DB.First(&user, "email = ?", email)
	return
}

func GetUsersByEmail(search string) (users []User) {

	DB.Find(&users, "email LIKE ?", "%"+search+"%")
	return
}
