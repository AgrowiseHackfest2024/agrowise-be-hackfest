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

func GetAllProductsHandler(ctx *fiber.Ctx) error {
	var product []entity.Product
	database.DB.Find(&product)
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Product data retrieved successfully",
		"data":    product,
	})
}

func GetProductByIDHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var product entity.Product

	result := database.DB.Preload("RatingProduct").Where("id = ?", id).First(&product)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	} else if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error retrieving Product data",
			"error":   result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Product data retrieved successfully",
		"data":    product,
	})
}
