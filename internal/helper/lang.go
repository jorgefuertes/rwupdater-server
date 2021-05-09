package helper

import (
	"html/template"

	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

type Lang struct {
	Abbr   string
	Link   string
	Active bool
}

func (h *Helper) langLink(c *fiber.Ctx, name string) string {
	return "/front/" + name + "/" + c.Params("color")
}

// SetPageTile - Set the current page title
func (h *Helper) SetPageTitle(key string) {
	h.PageTitle = h.T(key)
	h.IsMain = (key == "menu.main.title")
}

// T - Translate
func (h *Helper) T(key string, arg ...interface{}) template.HTML {
	return locale.T(h.Lang, key, arg...)
}

func (h *Helper) fillLangs(c *fiber.Ctx) {
	h.Lang = locale.GetUserLang(c)
	h.Langs = make([]Lang, 0)
	h.Langs = append(h.Langs,
		Lang{"en", h.langLink(c, "en"), (h.Lang == "en")},
		Lang{"es", h.langLink(c, "es"), (h.Lang == "es")},
	)
}

func (h *Helper) IsActiveLang(name string) bool {
	return (name == h.Lang)
}
