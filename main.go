package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "/public")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "usuario desde el backend",
		})
	})

	app.Listen(":4000")
	fmt.Println("Servidor corriendo en puerto 4000")
}
