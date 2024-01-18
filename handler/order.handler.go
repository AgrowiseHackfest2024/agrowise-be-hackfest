package handler

import (
	"agrowise-be-hackfest/database"
	"errors"
	"fmt"
	"os"
	"sync"

	// "time"

	"agrowise-be-hackfest/model/dto"
	"agrowise-be-hackfest/model/entity"
	"agrowise-be-hackfest/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func GetUserOrderHistory(ctx *fiber.Ctx) error {
	id := ctx.Locals("id")
	var order []entity.Order

	result := database.DB.Preload("Farmer").Preload("OrderItem.Product").Where("user_id = ?", id).Find(&order)

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
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	orderRequest := new(dto.OrderRequestDTO)
	if err := ctx.BodyParser(orderRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing order request",
			"error":   err.Error(),
		})
	}
	fmt.Println(orderRequest)

	validate := validator.New()
	errValidate := validate.Struct(orderRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error validating order request",
			"error":   errValidate.Error(),
		})
	}

	var s snap.Client
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	orderId := uuid.New()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId.String(),
			GrossAmt: int64((orderRequest.Price * orderRequest.Quantity) + orderRequest.ProductionFee),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    orderRequest.ProductID.String(),
				Name:  orderRequest.Name,
				Price: int64(orderRequest.Price),
				Qty:   int32(orderRequest.Quantity),
			},
			{
				ID:    orderRequest.FarmerID.String(),
				Name:  "Production Fee",
				Price: int64(orderRequest.ProductionFee),
				Qty:   1,
			},
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan error, 2)

	go utils.CreateOrder(orderRequest, orderId, snapResp, &wg, ctx, ch)
	go utils.CreateOrderItem(orderRequest, orderId, snapResp, &wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var errors []error
	for err := range ch {
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error creating order and order item",
			"errors":  errors,
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"message": "Order data retrieved successfully",
		"data":    snapResp,
	})
}

func UpdateTransactionStatus(ctx *fiber.Ctx, id string, status entity.StatusPembayaran, paymentMethod string) error {
	var order entity.Order

	result := database.DB.Preload("OrderItem.Product").Where("id = ?", id).First(&order)
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

	order.Status = status
	order.PaymentMethod = paymentMethod

	result = database.DB.Save(&order)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Error updating order data",
			"error":   result.Error.Error(),
		})
	}

	if status == entity.Success {
		for _, orderItem := range order.OrderItem {
			product := orderItem.Product

			product.Stok -= orderItem.Quantity
			product.Sold += orderItem.Quantity

			result = database.DB.Save(&product)
			if result.Error != nil {
				return ctx.Status(500).JSON(fiber.Map{
					"message": "Error updating product data",
					"error":   result.Error.Error(),
				})
			}
		}
	}

	return nil
}

func OrderNotificationHandler(ctx *fiber.Ctx) error {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	var notificationPayload map[string]interface{}

	error := ctx.BodyParser(&notificationPayload)
	if error != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error decoding notification payload",
		})
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order ID not found in the notification payload",
		})
	}

	var c coreapi.Client
	c.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	transactionStatusResp, e := c.CheckTransaction(orderId)
	if e != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": e.GetMessage(),
		})
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "accept" {
					UpdateTransactionStatus(ctx, orderId, entity.Success, transactionStatusResp.PaymentType)
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				if transactionStatusResp.FraudStatus == "accept" {
					UpdateTransactionStatus(ctx, orderId, entity.Success, transactionStatusResp.PaymentType)
				}
			} else if transactionStatusResp.TransactionStatus == "deny" {
				UpdateTransactionStatus(ctx, orderId, entity.Failed, transactionStatusResp.PaymentType)
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				UpdateTransactionStatus(ctx, orderId, entity.Failed, transactionStatusResp.PaymentType)
			} else if transactionStatusResp.TransactionStatus == "pending" {
				UpdateTransactionStatus(ctx, orderId, entity.Pending, transactionStatusResp.PaymentType)
			}
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func GetLastOrderHandler(ctx *fiber.Ctx) error {
	id := ctx.Locals("id")
	var LastOrderFarmerResponse dto.LastOrderFarmerResponseDTO

	result := database.DB.Model(&entity.Order{}).
		Select("orders.farmer_id", "farmers.nama").
		Joins("JOIN farmers ON orders.farmer_id = farmers.id").
		Where("orders.user_id = ?", id).
		Order("orders.created_at DESC").
		Limit(1).
		Scan(&LastOrderFarmerResponse)

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
		"data":    LastOrderFarmerResponse,
	})
}
