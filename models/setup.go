package models

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var database *gorm.DB
	var err error
	if gin.Mode() == "debug" {
		database, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		database, err = gorm.Open(sqlite.Open("xsstower.db"), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&XssHit{}, &User{}, &Handler{}, &CollectedPage{}, &HandlerRbac{})
	if err != nil {
		return
	}

	DB = database
}
