package handler

import (
	"agrowise-be-hackfest/database"
	"agrowise-be-hackfest/model/dto"
	"agrowise-be-hackfest/model/entity"

	// "agrowise-be-hackfest/utils"
	// "time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddRatingFarmerHandler(ctx *fiber.Ctx) error {
	userId := ctx.Locals("id")

	farmerId := ctx.Params("farmer_id")
	if farmerId == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Farmer id is required",
		})
	}

	ratingFarmerRequest := new(dto.RatingFarmerRequestDTO)

	if err := ctx.BodyParser(ratingFarmerRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing rating farmer request",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(ratingFarmerRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error validating rating farmer request",
			"error":   errValidate.Error(),
		})
	}

	farmerUuid, err := uuid.Parse(farmerId)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing farmer ID",
			"error":   err.Error(),
		})
	}

	userUuid, err := uuid.Parse(userId.(string))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing user ID",
			"error":   err.Error(),
		})
	}

	ratingFarmer := entity.RatingFarmer{
		FarmerID: farmerUuid,
		UserID:   userUuid,
		Rating:   ratingFarmerRequest.Rating,
		Review:   ratingFarmerRequest.Review,
	}

	result := database.DB.Create(&ratingFarmer)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error saving rating to the database",
			"error":   result.Error.Error(),
		})
	}

	ratingFarmerResponse := dto.RatingFarmerResponseDTO{
		ID:       ratingFarmer.ID.String(),
		FarmerID: ratingFarmer.FarmerID.String(),
		Rating:   ratingFarmer.Rating,
		Review:   ratingFarmer.Review,
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Rating submitted successfully",
		"data":    ratingFarmerResponse,
	})
}

func AddRatingProductHandler(ctx *fiber.Ctx) error {
	userId := ctx.Locals("id")

	productId := ctx.Params("product_id")
	if productId == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Product id is required",
		})
	}

	ratingProductRequest := new(dto.RatingProductRequestDTO)

	if err := ctx.BodyParser(ratingProductRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing rating product request",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(ratingProductRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error validating rating product request",
			"error":   errValidate.Error(),
		})
	}

	productUuid, err := uuid.Parse(productId)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing product ID",
			"error":   err.Error(),
		})
	}

	userUuid, err := uuid.Parse(userId.(string))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing user ID",
			"error":   err.Error(),
		})
	}

	ratingProduct := entity.RatingProduct{
		ProductID: productUuid,
		UserID:    userUuid,
		Rating:    ratingProductRequest.Rating,
		Review:    ratingProductRequest.Review,
	}

	result := database.DB.Create(&ratingProduct)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error saving rating to the database",
			"error":   result.Error.Error(),
		})
	}

	ratingProductResponse := dto.RatingProductResponseDTO{
		ID:        ratingProduct.ID.String(),
		ProductID: ratingProduct.ProductID.String(),
		Rating:    ratingProduct.Rating,
		Review:    ratingProduct.Review,
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Rating submitted successfully",
		"data":    ratingProductResponse,
	})
}

func GetAllRatingFarmerHandler(ctx *fiber.Ctx) error {
	var ratingFarmer []entity.RatingFarmer
	database.DB.Preload("User").Preload("Farmer").Find(&ratingFarmer)
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Rating farmer data retrieved successfully",
		"data":    ratingFarmer,
	})
}

func GetAllRatingProductHandler(ctx *fiber.Ctx) error {
	var ratingProduct []entity.RatingProduct
	database.DB.Preload("User").Preload("Product").Find(&ratingProduct)
	return ctx.Status(200).JSON(fiber.Map{
		"message": "Rating product data retrieved successfully",
		"data":    ratingProduct,
	})
}
