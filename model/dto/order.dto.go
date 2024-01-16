package dto

import "github.com/google/uuid"

type OrderRequestDTO struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Price     int       `json:"price" validate:"required,numeric,min=1"`
	Quantity  int       `json:"quantity" validate:"required,numeric,min=1"`
}
