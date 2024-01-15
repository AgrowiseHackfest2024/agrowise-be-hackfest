package handler

import (
	"agrowise-be-hackfest/database"
	"errors"

	// "agrowise-be-hackfest/model/dto"
	"agrowise-be-hackfest/model/entity"
	// "agrowise-be-hackfest/utils"
	// "time"

	// "github.com/dgrijalva/jwt-go/v4"
	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllFarmersHandler(ctx *fiber.Ctx) error {
	var farmer []entity.Farmer
	database.DB.Find(&farmer)
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Farmer data retrieved successfully",
		"data":    farmer,
	})
}

func GetFarmerByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var farmer entity.Farmer

	result := database.DB.Preload("RatingFarmer").Where("id = ?", id).First(&farmer)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Farmer not found",
		})
	} else if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error retrieving farmer data",
			"error":   result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Farmer data retrieved successfully",
		"data":    farmer,
	})
}
