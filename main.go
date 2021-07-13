package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sidmohanty11/auth-jwt-golang/routes"
)

const PORT = ":8000"

func main() {
	// server connection
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	})) //cross-origin-resource-sharing
	app.Use(logger.New())

	routes.Setup(app)

	// server listening port
	app.Listen(PORT)

	log.Printf("Listening at PORT%s", PORT)
}
