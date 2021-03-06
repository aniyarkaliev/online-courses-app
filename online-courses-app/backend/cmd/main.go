package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"online-courses-app/database"
	"online-courses-app/routes"
)

func main() {
	//DB connect
	database.Connect()

	//Creating application
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//Setup all routes
	routes.Setup(app)

	//Running server
	app.Listen(":8000")
}


