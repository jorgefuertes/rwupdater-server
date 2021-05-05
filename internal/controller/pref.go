package controller

import (
	"regexp"

	"git.martianoids.com/queru/retroserver/internal/cfg"
	"git.martianoids.com/queru/retroserver/internal/helper"
	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

func prefCtrl(app *fiber.App) {
	g := app.Group("/pref")

	// change user's language
	g.Get("/lang/:lang", func(c *fiber.Ctx) error {
		matched, _ := regexp.MatchString(cfg.AvailableLangReg, c.Params("lang"))
		if matched {
			locale.SetUserLang(c, c.Params("lang"))
		} else {
			locale.SetUserLang(c, "en")
		}
		return c.Redirect(string(c.Context().Referer()), 302)
	})

	// change user's color
	g.Get("/color/:color", func(c *fiber.Ctx) error {
		h := helper.New(c)
		h.SetUserColor(c.Params("color"))
		return c.Redirect(string(c.Context().Referer()), 302)
	})
}
