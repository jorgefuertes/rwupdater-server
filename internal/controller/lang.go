package controller

import (
	"regexp"

	"git.martianoids.com/queru/retroserver/internal/cfg"
	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

func langCtrl(app *fiber.App) {
	g := app.Group("/lang")
	// change user's locale
	g.Get("/:lang", func(c *fiber.Ctx) error {
		matched, _ := regexp.MatchString(cfg.AvailableLangReg, c.Params("lang"))
		if matched {
			locale.SetUserLang(c, c.Params("lang"))
		} else {
			locale.SetUserLang(c, "en")
		}
		return c.Redirect(string(c.Context().Referer()), 302)
	})
}
