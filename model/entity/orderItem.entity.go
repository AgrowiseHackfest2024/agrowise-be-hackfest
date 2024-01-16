package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product   Product   `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price     int       `json:"price" gorm:"type:integer"`
	Quantity  int       `json:"quantity" gorm:"type:integer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	OrderID uuid.UUID `json:"order_id" gorm:"type:uuid"`
}
