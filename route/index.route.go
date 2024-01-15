package route

import (
	"agrowise-be-hackfest/handler"
	"agrowise-be-hackfest/handler/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,Authorization",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/login", handler.AuthHandlerLogin)

	// Farmer
	app.Get("/farmers", middleware.Auth, handler.GetAllFarmersHandler)
	app.Get("/farmers/:id", middleware.Auth, handler.GetFarmerByIDHandler)

	// Product
	app.Get("/products", middleware.Auth, handler.GetAllProductsHandler)
	app.Get("/products/:id", middleware.Auth, handler.GetProductByIDHandler)

	// Order
	app.Get("/orders", middleware.Auth, handler.GetUserOrderHistory)
}
