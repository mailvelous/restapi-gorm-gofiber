package migration

import (
	"restapi-gorm-gofiber/database"
	"restapi-gorm-gofiber/model/entity"
)

func RunMigration() {
	database.DB.AutoMigrate(&entity.Users{})
}
