package helper

import (
	"html/template"

	"git.martianoids.com/queru/retroserver/internal/locale"
	"github.com/gofiber/fiber/v2"
)

// Download
type Download struct {
	File    string
	Version string
	Os      string
	CPU     string
	Ext     string
	Size    string
}

// Link to download
func (d *Download) Link() string {
	return "/download/dist/" + d.File
}

// Helper
type Helper struct {
	Ctx       *fiber.Ctx
	Lang      string
	PageTitle template.HTML
	Vars      fiber.Map
	IsMain    bool
	Downloads []Download
	Latest    string
}

// New - Create and return new Helper
func New(c *fiber.Ctx) Helper {
	h := new(Helper)
	h.Lang = locale.GetUserLang(c)
	h.Vars = make(fiber.Map)
	h.Ctx = c
	return *h
}

// SetPageTile - Set the current page title
func (h *Helper) SetPageTitle(key string) {
	h.PageTitle = h.T(key)
	h.IsMain = (key == "menu.main")
}

// T - Translate
func (h *Helper) T(key string, arg ...interface{}) template.HTML {
	return locale.T(h.Lang, key, arg...)
}

// IsLang - True if Lang == param
func (h *Helper) IsLang(l string) bool {
	return (h.Lang == l)
}
