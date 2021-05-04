package controller

import (
	"git.martianoids.com/queru/retroserver/internal/catalog"
	"github.com/gofiber/fiber/v2"
)

func apiArchCtrl(app *fiber.App) {
	a := app.Group("/api/arch")

	a.Get("/", func(c *fiber.Ctx) error {
		list, err := catalog.ArchList()
		if err != nil {
			return err
		}

		return c.JSON(list)
	})
}
