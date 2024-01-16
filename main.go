package main

import (
	"agrowise-be-hackfest/database"
	// "agrowise-be-hackfest/database/migration"
	"agrowise-be-hackfest/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	// migration.RunMigration()

	app := fiber.New()

	route.SetupRoutes(app)

	app.Listen(":8080")
}
