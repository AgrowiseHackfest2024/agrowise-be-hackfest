package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Farmer struct {
	ID         uuid.UUID      `json:"id" gorm:"primary_key;unique;type:uuid;default:uuid_generate_v4()"`
	Nama       string         `json:"nama" gorm:"type:varchar(255)"`
	Alamat     string         `json:"alamat" gorm:"type:varchar(255)"`
	LuasLahan  string         `json:"luas_lahan" gorm:"type:varchar(255)"`
	NoTelp     string         `json:"no_telp" gorm:"type:varchar(255)"`
	JenisSawah string         `json:"jenis_sawah" gorm:"type:varchar(255)"`
	Foto       pq.StringArray `gorm:"type:varchar(255)[]" json:"foto"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`

	RatingFarmer []RatingFarmer `gorm:"foreignKey:FarmerID"`
	Products     []Product      `gorm:"many2many:farmer_products;"`
	Orders       []Order        `gorm:"foreignKey:FarmerID"`
}
