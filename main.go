package main

import (
	"github.com/Chaksack/staples-backend/database"
	"github.com/Chaksack/staples-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	//new addition
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
