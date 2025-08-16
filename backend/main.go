package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Trackify API 🚀")
	})

	// Health endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Start server on port 8080
	app.Listen(":8080")
}
