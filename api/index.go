package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World route by test 👋!")
	})

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	app.Listen(":")
}
