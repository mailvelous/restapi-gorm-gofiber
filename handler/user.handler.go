package handler

import (
	"log"
	"restapi-gorm-gofiber/database"
	"restapi-gorm-gofiber/model/entity"
	"github.com/gofiber/fiber/v2"
)


func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.Users
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(users)


	return nil
}

func UserHandlerCreate(c *fiber.Ctx) error {
	user := new(entity.Users)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	newUser := entity.Users{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error

	if errCreateUser!= nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", 
			"message": "Could not create user", 
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully created",
		"data": newUser,
	})
}