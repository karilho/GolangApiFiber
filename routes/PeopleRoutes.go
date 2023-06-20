package routes

import (
	"GolangApiFiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(c *fiber.App) {
	c.Get("/", handlers.Home)

	c.Get("/peoples", handlers.ListAll)

	c.Post("/new", handlers.CreateFact)

}
