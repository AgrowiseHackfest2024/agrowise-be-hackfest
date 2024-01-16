package utils

import (
	"agrowise-be-hackfest/database"
	"agrowise-be-hackfest/model/dto"
	"agrowise-be-hackfest/model/entity"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateOrderItem(orderRequest *dto.OrderRequestDTO, orderId uuid.UUID, snapResp *snap.Response, wg *sync.WaitGroup, ch chan<- error) {
	defer wg.Done()

	orderItemData := entity.OrderItem{
		ID:        orderId,
		ProductID: orderRequest.ProductID,
		Price:     orderRequest.Price,
		Quantity:  orderRequest.Quantity,
		OrderID:   orderId,
	}

	result := database.DB.Create(&orderItemData)
	if result.Error != nil {
		ch <- result.Error
		return
	}

	ch <- nil
}

func CreateOrder(orderRequest *dto.OrderRequestDTO, orderId uuid.UUID, snapResp *snap.Response, wg *sync.WaitGroup, ctx *fiber.Ctx, ch chan<- error) {
	defer wg.Done()

	orderData := entity.Order{
		ID:              orderId,
		UserID:          uuid.MustParse(ctx.Locals("id").(string)),
		Total:           int(orderRequest.Price * orderRequest.Quantity),
		Status:          "pending",
		SnapToken:       snapResp.Token,
		SnapRedirectUrl: snapResp.RedirectURL,
	}

	result := database.DB.Create(&orderData)
	if result.Error != nil {
		ch <- result.Error
		return
	}

	ch <- nil
}
