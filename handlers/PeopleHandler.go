package handlers

import (
	"GolangApiFiber/database"
	"GolangApiFiber/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Bem vindo ao APP")
}

func ListAll(c *fiber.Ctx) error {
	peoples := []models.People{}
	database.DB.Db.Find(&peoples)
	c.Set("Content-Type", "application/json")
	return c.Status(200).JSON(peoples)
}

func CreateFact(c *fiber.Ctx) error {
	people := new(models.Fact)
	err := c.BodyParser(&people)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Db.Create(&people)

	return c.Status(200).JSON(people)
}
