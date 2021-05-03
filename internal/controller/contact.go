package controller

import (
	"git.martianoids.com/queru/retroserver/internal/contact"
	"github.com/gofiber/fiber/v2"
)

func contactForm(app *fiber.App) {
	g := app.Group("/contact")

	g.Get("/", func(c *fiber.Ctx) error {
		return c.Render("contact/form", fiber.Map{
			"PageTitle": "Contact Form",
		}, "layouts/main")
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

		return c.Render("contact/sent", fiber.Map{
			"PageTitle": "Contact Form",
		}, "layouts/main")
	})
}
