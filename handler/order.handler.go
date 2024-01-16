package handler

import (
	"agrowise-be-hackfest/database"
	"errors"

	"agrowise-be-hackfest/model/dto"
	"agrowise-be-hackfest/model/entity"

	// "agrowise-be-hackfest/utils"
	// "time"

	// "github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserOrderHistory(ctx *fiber.Ctx) error {
	id := ctx.Locals("id")
	var order []entity.Order

	result := database.DB.Preload("Product").Where("user_id = ?", id).Find(&order)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Order not found",
		})
	} else if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error retrieving order data",
			"error":   result.Error.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Order data retrieved successfully",
		"data":    order,
	})
}

func AddOrderHandler(ctx *fiber.Ctx) error {
	id := ctx.Locals("id")
	orderRequest := new(dto.OrderRequestDTO)

	if err := ctx.BodyParser(orderRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing order request",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(orderRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error validating order request",
			"error":   errValidate.Error(),
		})
	}

	userId, err := uuid.Parse(id.(string))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing user ID",
			"error":   err.Error(),
		})
	}

	newOrder := entity.Order{
		UserID:     userId,
		ProductID:  orderRequest.ProductID,
		Quantity:   orderRequest.Quantity,
		TotalPrice: orderRequest.TotalPrice,
		Status:     "pending",
	}

	result := database.DB.Create(&newOrder)

	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error creating order",
			"error":   result.Error.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Order created successfully",
		"data":    orderRequest,
	})
}
