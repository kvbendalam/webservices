package handlers

import (
	"fmt"

	"github.com/divrhino/divrhino-trivia/models"
	"github.com/gofiber/fiber/v2"
	"github.com/kvbendalam/webservices/database"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&fact)
	return c.Status(200).JSON(fact)
}

func GetFactById(c *fiber.Ctx) error {
	id := c.Params("id")
	var facts models.Fact

	result := database.DB.Db.Find(&facts, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&facts)
}

func DeleteFact(c *fiber.Ctx) error {
	id := c.Params("id")
	var fact models.Fact

	if result := database.DB.Db.First(&fact, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	database.DB.Db.Delete(&fact)

	return c.Status(200).JSON(&fact)
}
