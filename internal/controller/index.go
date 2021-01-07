package controller

import (
	"git.martianoids.com/queru/retroserver/internal/banner"
	"github.com/gofiber/fiber/v2"
)

func index(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(banner.Title)
	})
}
