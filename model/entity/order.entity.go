package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusPembayaran string

const (
	Pending StatusPembayaran = "pending"
	Success StatusPembayaran = "success"
	Failed  StatusPembayaran = "failed"
)

type Order struct {
	ID              uuid.UUID        `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	UserID          uuid.UUID        `json:"user_id" gorm:"type:uuid"`
	User            User             `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FarmerID        uuid.UUID        `json:"farmer_id" gorm:"type:uuid"`
	Farmer          Farmer           `json:"farmer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total           int              `json:"total" gorm:"type:integer"`
	Status          StatusPembayaran `json:"status" gorm:"type:varchar(255)"`
	SnapToken       string           `json:"snap_token,omitempty" gorm:"type:varchar(255)"`
	SnapRedirectUrl string           `json:"snap_redirect_url,omitempty" gorm:"type:varchar(255)"`
	PaymentMethod   string           `json:"payment_method,omitempty" gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	OrderItem []OrderItem `gorm:"foreignKey:OrderID"`
}
