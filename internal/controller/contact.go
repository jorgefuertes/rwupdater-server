package controller

import (
	"git.martianoids.com/queru/retroserver/internal/contact"
	"git.martianoids.com/queru/retroserver/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func contactCtrl(app *fiber.App) {
	g := app.Group("/contact")

	g.Get("/", func(c *fiber.Ctx) error {
		h := helper.New(c)
		h.SetPageTitle("menu.contact.title")
		return h.Render("contact/form")
	})

	g.Post("/", func(c *fiber.Ctx) error {
		m := contact.New()
		m.Name = c.FormValue("name")
		m.ReplyTo = c.FormValue("email")
		m.Subject = c.FormValue("subject")
		m.Message = c.FormValue("message")
		m.Agent = c.Get("User-Agent")
		m.Lang = c.Get("Accept-Language")
		m.IP = c.IP()
		go m.Send()

		h := helper.New(c)
		h.SetPageTitle("menu.contact.sent")
		return h.Render("contact/sent")
	})
}
