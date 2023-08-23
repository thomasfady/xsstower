package models

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/thomasfady/xsstower/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSqlite(dsn string) (database *gorm.DB, err error) {
	database, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return
}

func ConnectMysql(dsn string) (database *gorm.DB, err error) {
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func ConnectPostgres(dsn string) (database *gorm.DB, err error) {
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}

func ConnectSqlServer(dsn string) (database *gorm.DB, err error) {
	database, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	return
}

func ConnectDatabase() {
	var database *gorm.DB
	var err error
	if gin.Mode() == "debug" {
		database, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		dsn, ok := config.GetConfig("database.dsn").(string)
		if ok {
			switch config.GetConfig("database.type") {
			case "sqlite":
				database, err = ConnectSqlite(dsn)
			case "mysql":
				database, err = ConnectMysql(dsn)
			case "postgres":
				database, err = ConnectPostgres(dsn)
			case "sqlserver":
				database, err = ConnectSqlServer(dsn)
			default:
				database, err = gorm.Open(sqlite.Open("xsstower.db"), &gorm.Config{})
			}
		} else {
			database, err = gorm.Open(sqlite.Open("xsstower.db"), &gorm.Config{})
		}

	}

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&XssHit{}, &User{}, &Handler{}, &CollectedPage{}, &HandlerRbac{})
	if err != nil {
		panic(err)
	}

	DB = database
}
