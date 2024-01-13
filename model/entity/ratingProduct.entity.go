package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingProduct struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	ProductID uuid.UUID `json:"product_id" gorm:"type:uuid"`
	Product   Product   `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating    int       `json:"rating" gorm:"type:integer"`
	Review    string    `json:"review" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
