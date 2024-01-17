package migration

import (
	"agrowise-be-hackfest/database"
	"agrowise-be-hackfest/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	// err := database.DB.Migrator().DropTable(&entity.User{}, &entity.Product{}, &entity.Farmer{}, &entity.Order{}, &entity.OrderItem{}, &entity.RatingFarmer{}, &entity.RatingProduct{})
	err := database.DB.Migrator().DropTable(&entity.Product{}, &entity.Farmer{})
	if err != nil {
		log.Println("Error dropping table")
	}

	// err := database.DB.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Farmer{}, &entity.Order{}, &entity.OrderItem{}, &entity.RatingFarmer{}, &entity.RatingProduct{})
	err = database.DB.AutoMigrate(&entity.Product{}, &entity.Farmer{})
	if err != nil {
		log.Println("Error running migration")
	}

	fmt.Println("Migration run successfully")
}
