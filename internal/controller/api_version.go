package controller

import (
	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func apiVersionCtrl(app *fiber.App) {
	a := app.Group("/api/version")

	a.Get("/client", func(c *fiber.Ctx) error {
		h := helper.New(c)
		h.FillDownloads()

		return c.JSON(fiber.Map{"latest": h.Downloads[0].Version})
	})
}
