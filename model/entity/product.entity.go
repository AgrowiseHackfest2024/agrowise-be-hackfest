package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Nama      string    `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi string    `json:"deskripsi" gorm:"type:varchar(255)"`
	Harga     int       `json:"harga" gorm:"type:integer"`
	Stok      int       `json:"stok" gorm:"type:integer"`
	Foto      []string  `json:"foto" gorm:"type:varchar(255)[]"`
	Sold      int       `json:"sold" gorm:"type:integer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RatingProduct []RatingProduct `gorm:"foreignKey:ProductID"`
	Order         []Order         `gorm:"foreignKey:ProductID"`
}
