package controller

import (
	"github.com/gofiber/fiber/v2"
)

func errorCtrl(app *fiber.App) {
	g := app.Group("/error")
	// change user's locale
	g.Get("/:code", func(c *fiber.Ctx) error {
		if c.Params("code") == "404" {
			return fiber.NewError(fiber.ErrNotFound.Code, "Resource not found")
		}
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Internal server error")
	})
}
