package dto

import "github.com/google/uuid"

type OrderRequestDTO struct {
	ProductID  uuid.UUID `json:"product_id" validate:"required"`
	Quantity   int       `json:"quantity" validate:"required,numeric,min=1"`
	TotalPrice int       `json:"total_price" validate:"required,numeric,min=1"`
}
