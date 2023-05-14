package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Listen(":3000")
	fmt.Println("Servidor corriendo en puerto 3000")
}
