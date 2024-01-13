package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RatingFarmer struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	FarmerID  uuid.UUID `json:"farmer_id" gorm:"type:uuid"`
	Farmer    Farmer    `json:"farmer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating    int       `json:"rating" gorm:"type:integer"`
	Review    string    `json:"review" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
