package handler

import (
	"agrowise-be-hackfest/database"
	"agrowise-be-hackfest/model/dto"
	"agrowise-be-hackfest/model/entity"
	"agrowise-be-hackfest/utils"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthHandlerLogin(ctx *fiber.Ctx) error {
	loginRequest := new(dto.LoginRequestDTO)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing login request",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error validating login request",
			"error":   errValidate.Error(),
		})
	}

	var existingUser entity.User
	err := database.DB.Where("email = ?", loginRequest.Email).First(&existingUser).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	isValid := utils.CheckPassword(existingUser.Password, loginRequest.Password)
	if !isValid {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	claims := jwt.MapClaims{}
	claims["id"] = existingUser.ID
	claims["email"] = existingUser.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := utils.GenerateToken(&claims)

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error generating token",
			"error":   err.Error(),
		})
	}

	responseDTO := dto.LoginResponseDTO{
		Message: "Login successful",
		Token:   token,
	}

	return ctx.Status(200).JSON(responseDTO)
}

func GetUserProfileHandler(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(string)

	var existingUser entity.User
	err := database.DB.Where("id = ?", id).First(&existingUser).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "User not found",
			"error":   err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Success",
		"user":    existingUser,
	})
}
