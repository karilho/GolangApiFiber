package main

import (
	"GolangApiFiber/database"
	"GolangApiFiber/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	database.ConnectDb()
	fmt.Println("Aplicação iniciada")
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
