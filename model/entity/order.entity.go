package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User       User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID  uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product    Product   `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity   int       `json:"quantity" gorm:"type:integer"`
	TotalPrice int       `json:"total_price" gorm:"type:integer"`
	Status     string    `json:"status" gorm:"type:varchar(255)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
