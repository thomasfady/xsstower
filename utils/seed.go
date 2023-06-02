package utils

import (
	"log"

	"github.com/thomasfady/xsstower/models"

	"github.com/gin-gonic/gin"
)

func SeedAdmin() {
	var users []models.User
	models.DB.Where(models.User{IsAdmin: true}).Find(&users)
	if len(users) == 0 {
		var pass string
		if gin.Mode() == "debug" {
			pass = "test"
			admin := models.User{
				Email:    "admin@admin.fr",
				Username: "admin",
				IsAdmin:  true,
			}
			admin.SetPassword(pass)
			models.DB.Create(&admin)
			user := models.User{
				Email:    "user@user.fr",
				Username: "user",
			}
			user.SetPassword(pass)
			models.DB.Create(&user)
			user2 := models.User{
				Email:    "user2@user.fr",
				Username: "user2",
			}
			user2.SetPassword(pass)
			models.DB.Create(&user2)
			rbac := models.HandlerRbac{
				User:       admin,
				Permission: models.OWNER,
			}
			rbac2 := models.HandlerRbac{
				User:       user,
				Permission: models.READ,
			}
			handler := models.Handler{
				Path:           "admin",
				OwnerID:        int(admin.ID),
				Screenshot:     true,
				Dom:            true,
				CollectedPages: []string{"login"},
				Members:        []models.HandlerRbac{rbac, rbac2},
			}
			models.DB.Create(&handler)
		} else {
			pass = RandomHex(10)
			admin := models.User{
				Email:    "admin@admin.fr",
				Username: "admin",
				IsAdmin:  true,
			}
			admin.SetPassword(pass)
			models.DB.Create(&admin)
		}

		

		log.Printf("New admin account created with email '%s' and password '%s'", "admin@admin.fr", pass)

	}
}
