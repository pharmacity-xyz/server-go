package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pharmacity-xyz/server/handlers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	api.Post("/createuser", handlers.CreateUser)
}

func main() {
	app := fiber.New()

	Routes(app)

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Unable to start server")
	}
}
