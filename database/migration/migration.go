package migration

import (
	"agrowise-be-hackfest/database"
	"agrowise-be-hackfest/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Farmer{}, &entity.Product{}, &entity.Order{}, &entity.RatingFarmer{}, &entity.RatingProduct{})

	if err != nil {
		log.Println("Error running migration")
	}

	fmt.Println("Migration run successfully")
}
