package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/roman-mik/horus-heresy-tactica/routes/api"
)

func main() {
	app := fiber.New()

	app.Get("/api/legions", api.GetLegions)
	app.Get("/api/legions/:id", api.GetLegionByID)

	app.Listen(":3000")
}
