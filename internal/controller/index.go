package controller

import (
	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func indexCtrl(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		h := helper.New(c)
		h.SetPageTitle("menu.main.title")
		return h.Render("index")
	})
}
