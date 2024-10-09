package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roman-mik/horus-heresy-tactica/models"
)

func GetLegions(c *fiber.Ctx) error {
	return c.JSON(models.GetLegions())
}

func GetLegionByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Invalid ID")
	}

	legion := models.GetLegionByID(id)
	if legion == nil {
		return c.Status(404).SendString("Legion not found")
	}

	return c.JSON(legion)
}
