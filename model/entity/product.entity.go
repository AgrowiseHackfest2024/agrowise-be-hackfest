package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Nama      string         `json:"nama" gorm:"type:varchar(255)"`
	Deskripsi string         `json:"deskripsi" gorm:"type:varchar(255)"`
	Stok      int            `json:"stok" gorm:"type:integer"`
	Harga     int            `json:"harga" gorm:"type:integer"`
	Foto      pq.StringArray `gorm:"type:varchar(255)[]" json:"foto"`
	Sold      int            `json:"sold" gorm:"type:integer"`
	Berat     int            `json:"berat" gorm:"type:integer"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	RatingProduct []RatingProduct `gorm:"foreignKey:ProductID"`
	OrderItem     []OrderItem     `gorm:"foreignKey:ProductID"`
}
