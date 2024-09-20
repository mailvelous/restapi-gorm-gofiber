package handler

import (
	"github.com/gofiber/fiber/v2"
	"restapi-gorm-gofiber/model/request"
	"github.com/go-playground/validator/v10"
	"restapi-gorm-gofiber/model/entity"
	"restapi-gorm-gofiber/database"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Validation error",
			"error": errValidate.Error(),
		})
	}

	var user entity.Users
	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "User not found",
			"data": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"token": "secret",
	})
}