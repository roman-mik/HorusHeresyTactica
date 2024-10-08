package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/roman-mik/horus-heresy-tactica/models"
)

// GetLegions handles GET requests for the list of legions
func GetLegions(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(models.GetLegions())
}

func main() {
	app := fiber.New()

	app.Get("/legions", GetLegions)

	app.Listen(":3000")
}
