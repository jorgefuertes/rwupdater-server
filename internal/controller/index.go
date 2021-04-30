package controller

import (
	"github.com/gofiber/fiber/v2"
)

func index(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"PageTitle": "Main Menu",
		}, "layouts/main")
	})
}
