package helper

import (
	"html/template"

	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

// Helper
type Helper struct {
	Ctx       *fiber.Ctx
	Lang      string
	PageTitle template.HTML
	Vars      fiber.Map
	IsMain    bool
	Downloads []Download
	Latest    string
	Err       string
	Colors    []Color
	Langs     []Lang
}

// New - Create and return new Helper
func New(c *fiber.Ctx) Helper {
	h := new(Helper)
	h.Ctx = c
	h.Colors = make([]Color, 0)
	h.Colors = append(h.Colors,
		Color{"W", h.IsActiveColor("W")},
		Color{"G", h.IsActiveColor("G")},
		Color{"B", h.IsActiveColor("B")},
		Color{"A", h.IsActiveColor("A")},
	)
	h.Lang = locale.GetUserLang(c)
	h.Langs = make([]Lang, 0)
	h.Langs = append(h.Langs, Lang{"en", h.IsActiveLang("en")}, Lang{"es", h.IsActiveLang("es")})
	h.Vars = make(fiber.Map)
	return *h
}
