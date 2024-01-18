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
	app.Post("/register", handler.AuthHandlerRegister)
	app.Get("/profile", middleware.Auth, handler.GetUserProfileHandler)

	// Farmer
	app.Get("/farmers", handler.GetAllFarmersHandler)
	app.Get("/farmers/:id", handler.GetFarmerByIDHandler)

	// Product
	app.Get("/products", handler.GetAllProductsHandler)
	app.Get("/products/:id", handler.GetProductByIDHandler)

	// Order
	app.Get("/orders", middleware.Auth, handler.GetUserOrderHistory)
	app.Get("/orders/last", middleware.Auth, handler.GetLastOrderHandler)
	app.Post("/orders", middleware.Auth, handler.AddOrderHandler)
	app.Post("/orders/notification", handler.OrderNotificationHandler)

	// Rating
	app.Get("/rating/farmer", middleware.Auth, handler.GetAllRatingFarmerHandler)
	app.Get("/rating/product", middleware.Auth, handler.GetAllRatingProductHandler)
	app.Post("/rating/farmer/:farmer_id", middleware.Auth, handler.AddRatingFarmerHandler)
	app.Post("/rating/product/:product_id", middleware.Auth, handler.AddRatingProductHandler)
}
