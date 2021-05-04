package controller

import (
	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

// document controller
func docCtrl(app *fiber.App) {
	g := app.Group("/doc")

	g.Get("/:doc", func(c *fiber.Ctx) error {
		h := helper.New(c)
		h.SetPageTitle("menu." + c.Params("doc") + ".title")
		return h.Render("doc/" + c.Params("doc"))
	})
}
