package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Nama      string    `json:"nama" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255)"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RatingFarmer  []RatingFarmer  `gorm:"foreignKey:UserID"`
	RatingProduct []RatingProduct `gorm:"foreignKey:UserID"`
	Order         []Order         `gorm:"foreignKey:UserID"`
}
